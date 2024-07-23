package svc

import (
	"easy-chat/apps/social/rpc/internal/config"
	"easy-chat/apps/social/socialModels"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	socialModels.FriendsModel
	socialModels.FriendRequestsModel
	socialModels.GroupsModel
	socialModels.GroupRequestsModel
	socialModels.GroupMembersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		FriendsModel:        socialModels.NewFriendsModel(sqlConn, c.Cache),
		FriendRequestsModel: socialModels.NewFriendRequestsModel(sqlConn, c.Cache),
		GroupsModel:         socialModels.NewGroupsModel(sqlConn, c.Cache),
		GroupRequestsModel:  socialModels.NewGroupRequestsModel(sqlConn, c.Cache),
		GroupMembersModel:   socialModels.NewGroupMembersModel(sqlConn, c.Cache),
	}
}
