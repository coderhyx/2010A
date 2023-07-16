package logic

import (
	"context"

	"goods_srv/internal/svc"
	"goods_srv/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelGoodsLogic {
	return &DelGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelGoodsLogic) DelGoods(in *goods.DelGoodsReq) (*goods.DelGoodsResp, error) {
	// todo: add your logic here and delete this line

	return &goods.DelGoodsResp{}, nil
}
