package logic

import (
	"context"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type User struct {
	Id    string
	Name  string
	Phone string
	Pass  string
}

var (
	users = map[string]*User{
		"1": {
			Id:    "1",
			Name:  "张三",
			Phone: "13000001111",
			Pass:  "123456",
		},
		"2": {
			Id:    "2",
			Name:  "李四",
			Phone: "132000111",
			Pass:  "123456",
		},
	}
)

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	if u, ok := users[in.Id]; ok {
		return &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}, nil
	}

	return &user.GetUserResp{}, nil
}
