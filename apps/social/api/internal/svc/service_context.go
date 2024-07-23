package svc

import (
	"easy-chat/apps/social/api/internal/config"
	"easy-chat/apps/social/rpc/social_client"
	"easy-chat/apps/user/rpc/user_client"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc   user_client.User
	SocialRpc social_client.Social
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		UserRpc:   user_client.NewUser(zrpc.MustNewClient(c.UserRpc)),
		SocialRpc: social_client.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
	}
}
