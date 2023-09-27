# 构建环境
FROM golang:alpine AS builder
# 工作目录
WORKDIR /go/app/
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
# 单独层下载golang依赖包
COPY ./go.mod /go/app/go.mod
COPY ./go.sum /go/app/go.sum
RUN go mod download
COPY ./ /go/app/
RUN go mod vendor
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./cmd/main.go ./cmd/wire_gen.go

# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM alpine AS runner
# 全局工作目录
WORKDIR /go/app
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /go/app/main .
# 复制编译阶段里的config文件夹到目标目录
COPY --from=builder /go/app/config/config.yml .
# 环境变量
# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata
CMD ["./main", "-config", "/go/app/config.yml"]
