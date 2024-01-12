# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.19
# 设置工作目录
WORKDIR /data/golang/my-gin-demo
# 将本地文件复制到容器中
COPY . .
# 使用 Go Modules 下载依赖  编译项目
RUN go env -w GO111MODULE=on \
   && go env -w GOPROXY=https://goproxy.cn,direct \
   && go env -w CGO_ENABLED=0 \
   && go mod download \
   && go mod tidy \
   && go build -o server .

# 最终镜像使用轻量的 alpine 镜像
FROM alpine:latest
# 添加作者
LABEL MAINTAINER="baily@gmail.com"
# 设置工作目录
WORKDIR /data/golang/my-gin-demo
# 将二进制文件从前一个镜像中复制到这里
COPY --from=0 /data/golang/my-gin-demo/server ./
# 暴露端口
EXPOSE 8083
# 启动应用程序
ENTRYPOINT ./server
