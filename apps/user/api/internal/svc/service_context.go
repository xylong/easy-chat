package svc

import (
	"easy-chat/apps/user/api/internal/config"
	"easy-chat/apps/user/api/internal/middleware"
	"easy-chat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	UserClient        userclient.User
	LoginVerification rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserClient:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle,
	}
}
