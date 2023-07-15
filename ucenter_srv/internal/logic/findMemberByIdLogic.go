package logic

import (
	"context"
	"fmt"

	"ucenter_srv/internal/svc"
	"ucenter_srv/pb/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMemberByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMemberByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMemberByIdLogic {
	return &FindMemberByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 会员查找
func (l *FindMemberByIdLogic) FindMemberById(in *ucenter.MemberReq) (*ucenter.MemberInfo, error) {
	// todo: add your logic here and delete this line
	m, _ := l.svcCtx.MemberModel.FindOne(l.ctx, in.MemberId)
	fmt.Println(m)
	return &ucenter.MemberInfo{}, nil
}
