package logic

import (
	"context"
	"ucenter_srv/internal/model"

	"ucenter_srv/internal/svc"
	"ucenter_srv/pb/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByPhoneLogic {
	return &RegisterByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterByPhoneLogic) RegisterByPhone(in *ucenter.RegReq) (*ucenter.RegRes, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.MemberModel.Insert(l.ctx, &model.Member{})
	return &ucenter.RegRes{}, nil
}
