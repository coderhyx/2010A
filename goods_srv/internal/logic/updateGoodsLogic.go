package logic

import (
	"context"

	"goods_srv/internal/svc"
	"goods_srv/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsLogic {
	return &UpdateGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGoodsLogic) UpdateGoods(in *goods.UpdateGoodsReq) (*goods.UpdateGoodsResp, error) {
	// todo: add your logic here and delete this line

	return &goods.UpdateGoodsResp{}, nil
}
