package user

import (
	"context"

	"ucenter_api/internal/svc"
	"ucenter_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeRequest) (resp *types.BoolResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
