package logic

import (
	"context"
	"easy-chat/apps/user/models"
	"easy-chat/pkg/sqlx"
	"time"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateReq) (*user.Response, error) {
	now := time.Now()
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &models.Users{
		Avatar:    in.Avatar,
		Name:      in.Name,
		Phone:     in.Phone,
		Password:  sqlx.ToNullString(in.Password),
		CreatedAt: sqlx.ToNullTime(now),
		UpdatedAt: sqlx.ToNullTime(now),
	})

	return &user.Response{}, err
}
