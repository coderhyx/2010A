package user

import (
	common "2010A/common/ctl"
	"context"
	"ucenter_srv/ucenterclient"

	"github.com/jinzhu/copier"

	"ucenter_api/internal/svc"
	"ucenter_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MembersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MembersLogic {
	return &MembersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MembersLogic) Members(req *types.GetMembersRequest) (res *common.Result, err error) {
	mems, _ := l.svcCtx.UCenterSrv.FindMembers(l.ctx, &ucenterclient.MembersReq{
		Page:     req.Page,
		Limit:    req.PageSize,
		UserName: req.UserName,
	})
	resObj := common.NewResult()
	resp := &types.GetMembersResponse{}
	copier.Copy(resp, mems)
	res = resObj.Deal(resp, err)
	return res, nil
}
