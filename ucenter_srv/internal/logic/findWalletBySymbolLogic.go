package logic

import (
	"context"

	"ucenter_srv/internal/svc"
	"ucenter_srv/pb/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindWalletBySymbolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindWalletBySymbolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindWalletBySymbolLogic {
	return &FindWalletBySymbolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindWalletBySymbolLogic) FindWalletBySymbol(in *ucenter.AssetReq) (*ucenter.MemberWallet, error) {
	// todo: add your logic here and delete this line

	return &ucenter.MemberWallet{}, nil
}
