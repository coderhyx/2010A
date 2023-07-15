package user

import (
	common "2010A/common/ctl"
	"context"
	"ucenter_srv/pb/ucenter"

	"github.com/jinzhu/copier"

	"ucenter_api/internal/svc"
	"ucenter_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneLogic {
	return &FindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOneLogic) FindOne(req *types.FindOneRequest) (res *common.Result, err error) {
	m, err := l.svcCtx.UCenterSrv.FindMemberById(l.ctx, &ucenter.MemberReq{
		MemberId: req.MemberId,
	})
	resp := &types.FindOneResponse{}
	copier.Copy(resp, m)
	resObj := common.NewResult()
	res = resObj.Deal(resp, err)
	return res, nil
}
