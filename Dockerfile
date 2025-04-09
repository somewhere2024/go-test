# 使用官方的 Golang 镜像作为构建阶段的基础镜像
# 这里选择了 Go 1.20 版本，你可以根据需要更改版本
FROM golang:1.24 AS build

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 文件复制到容器中
COPY go.mod go.sum ./

# 下载依赖项
RUN go mod tidy

# 将整个项目复制到容器中的工作目录
COPY . .

# 编译 Go 项目，生成二进制文件
RUN GOOS=linux GOARCH=amd64 go build -o myapp .

# 使用一个较小的基础镜像来运行应用（例如使用 alpine 镜像）
FROM alpine:latest

# 安装必要的依赖（例如 SSL 库）
RUN apk --no-cache add ca-certificates

# 将构建阶段生成的二进制文件复制到当前镜像中
COPY --from=build /app/myapp /app/myapp

# 设置容器启动时运行的命令
ENTRYPOINT ["/app/myapp"]

# 暴露应用运行的端口（例如，应用监听端口 8080）
EXPOSE 8080

# 设置容器的工作目录
WORKDIR /app

# 默认启动命令
CMD ["/app/myapp"]
