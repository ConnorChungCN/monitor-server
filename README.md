# 算法管理服务

# 目录结构
``` shell
.
├── README.md       说明文件
├── adapter         适配器层
├── app             应用层
├── cmd             启动文件
├── config          配置文件
├── dockerfile      镜像生成文件
├── domain          领域层
├── go.mod          
├── go.sum
├── infrastructure  基础设施层
├── pkg             其他包
├── scripts         脚本
├── template        算法开发模板
└── utils           工具方法
```

# 依赖工具
- mysql: 算法结构化数据存储
- minio: 算法镜像存储

# 命令
```shell
# 生成 grpc proto
protoc --proto_path=./adapter/grpc/proto  --proto_path=. --go_out=paths=source_relative:./adapter/grpc/proto --go-grpc_out=paths=source_relative:./adapter/grpc/proto {prootofile} 

# 本地启动
go run ./cmd/main.go ./cmd/wire_gen.go

# 编译
go build ./cmd/main.go ./cmd/wire_gen.go
```

