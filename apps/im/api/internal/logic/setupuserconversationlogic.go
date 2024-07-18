package logic

import (
	"context"
	"easy-chat/apps/im/rpc/imclient"

	"easy-chat/apps/im/api/internal/svc"
	"easy-chat/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUpUserConversationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 建立会话
func NewSetUpUserConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUpUserConversationLogic {
	return &SetUpUserConversationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUpUserConversationLogic) SetUpUserConversation(req *types.SetUpUserConversationReq) (
	resp *types.SetUpUserConversationResp, err error) {

	_, err = l.svcCtx.Im.SetUpUserConversation(l.ctx, &imclient.SetUpUserConversationReq{
		SendId:   req.SendId,
		RecvId:   req.RecvId,
		ChatType: req.ChatType,
	})

	return
}