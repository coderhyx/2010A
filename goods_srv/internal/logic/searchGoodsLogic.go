package logic

import (
	"context"

	"goods_srv/internal/svc"
	"goods_srv/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGoodsLogic {
	return &SearchGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchGoodsLogic) SearchGoods(in *goods.SearchGoodsReq) (*goods.SearchGoodsResp, error) {
	// todo: add your logic here and delete this line

	return &goods.SearchGoodsResp{}, nil
}
