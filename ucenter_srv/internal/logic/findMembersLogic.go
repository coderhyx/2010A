package logic

import (
	"context"

	"github.com/jinzhu/copier"

	"ucenter_srv/internal/svc"
	"ucenter_srv/pb/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMembersLogic {
	return &FindMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 会员列表
func (l *FindMembersLogic) FindMembers(in *ucenter.MembersReq) (*ucenter.MembersRes, error) {
	where := ""
	if in.GetUserName() != "" {
		where = "username =" + in.GetUserName()
	}
	members, _ := l.svcCtx.MemberModel.FindMembers(l.ctx, where, in.Page, in.Limit)
	rep := &ucenter.MembersRes{}
	rep.Total, _ = l.svcCtx.MemberModel.GetTotal(l.ctx, where)
	copier.Copy(&rep.Members, &members)
	return rep, nil
}
