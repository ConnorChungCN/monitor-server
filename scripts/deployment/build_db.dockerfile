FROM mysql:latest

WORKDIR /docker-entrypoint-initdb.d
ENV LANG=C.UTF-8

ADD ./scripts/db/init.sql .

RUN /docker-entrypoint-initdb.d