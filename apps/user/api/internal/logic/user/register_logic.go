package user

import (
	"context"
	"easy-chat/apps/user/rpc/user"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"

	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register 注册
func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	var (
		err  error
		resp types.RegisterResp
	)

	registerResp, err := l.svcCtx.UserClient.Register(l.ctx, &user.RegisterReq{
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Password: req.Password,
		Avatar:   req.Avatar,
		Sex:      cast.ToInt32(req.Sex),
	})
	if err != nil {
		return nil, err
	}

	copier.Copy(&resp, registerResp)
	return &resp, nil
}
