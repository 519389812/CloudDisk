# CloudDisk
> 轻量级云盘系统

使用到的库
```text
go get xorm.io/xorm

go install github.com/zeromicro/go-zero/tools/goctl@latest
goctl migrate —verbose —version v1.3.3

go get github.com/jordan-wright/email

go get github.com/go-redis/redis/v8

```

使用到的命令
```text
# 创建api服务
goctl api new core
# 启动服务
go run core.go -f etc/core-api.yaml
# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero
```