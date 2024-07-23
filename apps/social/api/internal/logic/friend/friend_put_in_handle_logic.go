package friend

import (
	"context"
	"easy-chat/apps/social/rpc/social_client"
	"easy-chat/pkg/ctxdata"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FriendPutInHandle 好友申请处理
func (l *FriendPutInHandleLogic) FriendPutInHandle(req *types.FriendPutInHandleReq) (resp *types.FriendPutInHandleResp, err error) {
	_, err = l.svcCtx.SocialRpc.FriendPutInHandle(l.ctx, &social_client.FriendPutInHandleReq{
		FriendReqId:  req.FriendReqId,
		UserId:       ctxdata.GetUId(l.ctx),
		HandleResult: req.HandleResult,
	})

	return
}
