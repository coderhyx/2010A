package logic

import (
	"context"

	"goods_srv/internal/svc"
	"goods_srv/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsByIdLogic {
	return &GetGoodsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsByIdLogic) GetGoodsById(in *goods.GetGoodsByIdReq) (*goods.GetGoodsByIdResp, error) {
	// todo: add your logic here and delete this line

	return &goods.GetGoodsByIdResp{}, nil
}
