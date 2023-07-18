package logic

import (
	"context"
	"fmt"
	"goods_srv/internal/dao"

	"goods_srv/internal/svc"
	"goods_srv/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	goodsDao dao.GoodsDao
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
	gm, _ := l.goodsDao.FindById(l.ctx, in.Id)
	fmt.Println("-------------->", gm)
	return &goods.GetGoodsByIdResp{}, nil
}
