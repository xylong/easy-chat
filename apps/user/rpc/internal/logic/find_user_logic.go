package logic

import (
	"context"
	"easy-chat/apps/user/models"
	"github.com/jinzhu/copier"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	var (
		err   error
		users []*models.Users
	)

	if in.Phone != "" {
		userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
		if err == nil {
			users = append(users, userEntity)
		} else if in.Name != "" {
			users, err = l.svcCtx.UsersModel.ListByName(l.ctx, in.Name)
		} else if len(in.Ids) > 0 {
			users, err = l.svcCtx.UsersModel.ListByIds(l.ctx, in.Ids)
		}
	}
	if err != nil {
		return nil, err
	}

	var resp []*user.UserEntity
	copier.Copy(&resp, users)
	return &user.FindUserResp{
		User: resp,
	}, err
}
