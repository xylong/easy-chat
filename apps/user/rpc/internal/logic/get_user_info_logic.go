package logic

import (
	"context"
	"easy-chat/apps/user/models"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var (
	UserNotFoundErr = errors.New("用户不存在")
)

// GetUserInfo 获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	userEntity, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, UserNotFoundErr
		}

		return nil, err
	}

	var resp user.UserEntity
	_ = copier.Copy(&resp, userEntity)

	return &user.GetUserInfoResp{
		User: &resp,
	}, err
}
