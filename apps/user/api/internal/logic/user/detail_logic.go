package user

import (
	"context"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/ctxdata"
	"github.com/jinzhu/copier"

	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Detail 用户详情
func (l *DetailLogic) Detail(req *types.UserInfoReq) (*types.UserInfoResp, error) {
	var (
		err  error
		resp types.User

		uid = ctxdata.GetUId(l.ctx)
	)

	userInfoResp, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	copier.Copy(&resp, userInfoResp.User)
	return &types.UserInfoResp{
		Info: resp,
	}, nil
}
