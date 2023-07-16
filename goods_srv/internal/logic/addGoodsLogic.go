package logic

import (
	"context"

	"goods_srv/internal/svc"
	"goods_srv/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodsLogic {
	return &AddGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------goods-----------------------
func (l *AddGoodsLogic) AddGoods(in *goods.AddGoodsReq) (*goods.AddGoodsResp, error) {
	// todo: add your logic here and delete this line

	return &goods.AddGoodsResp{}, nil
}
