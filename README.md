即时通讯

### 库
- go-zero

### 构建镜像容器
```dockerfile
docker build -t go:alpine3.18.4 .
docker run --name go -d go:alpine3.18.4

# 删除
docker stop go | xargs docker rm
```

### 目录
```
.
├── apps
│   ├── social                  社交服务
│   └── user                    用户服务
├── cmd.sh
├── deploy
│   └── sql
├── docker-compose.yaml
├── go.mod
├── go.sum
└── pkg
    ├── constants
    ├── ctxdata
    ├── encrypt
    ├── interceptor
    ├── resultx
    ├── sqlx
    ├── wuid
    └── xerr
```