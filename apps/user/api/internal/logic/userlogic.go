package logic

import (
	"context"
	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/apps/user/api/internal/types"
	"easy-chat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserResp, err error) {
	getUserResp, err := l.svcCtx.UserClient.GetUser(l.ctx, &userclient.GetUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserResp{
		Id:    getUserResp.Id,
		Name:  getUserResp.Name,
		Phone: getUserResp.Phone,
	}, nil
}
