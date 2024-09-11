即时通讯

### 库
- go-zero

### 构建镜像容器
```dockerfile
docker build -t go:alpine3.18.4 .
docker run --name go -d go:alpine3.18.4

# 删除
docker stop go | xargs docker rm

# etcd
docker build -t etcd -f ./Dockerfile-etcd .
docker run -d -p 2379:2379 -p 2380:2380 --net user --ip 168.10.0.20 --name etcd etcd

# 编译user
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/user-api ./api/user.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/user-rpc ./rpc/user.go

docker build -t user-api -f ./Dockerfile .
docker build -t user-rpc -f ./Dockerfile_rpc .

# 部署
docker run --name etcd -p 2379:2379 --net user --ip 168.10.0.20 -d etcd
docker run --name user-rpc -p 8080:8080 --net user --ip 168.10.0.70 --link etcd -d user-rpc
docker run --name user-api -p 8888:8888 --net user --ip 168.10.0.50 --link etcd -d user-api

# etcd查询所有的key
etcdctl get --prefix ""

docker ps

```



### 目录
```
.
├── Dockerfile
├── Makefile                            项目编译脚本程序
├── Dockerfile_etcd
├── Dockerfile_rpc
├── Makefile
├── README.md
├── apps
│   ├── im
│   ├── social
│   └── user
├── cmd.sh
├── bin                                 编译后的文件
├── deploy                              项目部署相关信息
│   ├── dockerfile
│   ├── mk
│   └── sql
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
└── pkg                                 公共工具目录
    ├── constants
    ├── ctxdata
    ├── encrypt
    ├── interceptor
    ├── resultx
    ├── sqlx
    ├── tools
    ├── wuid
    └── xerr
```