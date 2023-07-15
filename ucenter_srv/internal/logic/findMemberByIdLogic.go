package logic

import (
	"context"
	"fmt"
	"ucenter_srv/internal/svc"
	"ucenter_srv/pb/ucenter"

	"github.com/jinzhu/copier"
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
	m, e := l.svcCtx.MemberModel.FindOne(l.ctx, in.MemberId)
	if e != nil {
		return nil, e
	}
	resp := &ucenter.MemberInfo{}
	copier.Copy(resp, m)
	fmt.Println(m)
	return resp, nil
}
