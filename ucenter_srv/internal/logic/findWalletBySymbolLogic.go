package logic

import (
	"context"
	"fmt"
	"ucenter_srv/internal/svc"
	"ucenter_srv/pb/ucenter"

	"github.com/pkg/errors"
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
	e := e1()
	fmt.Printf("e1 is %+v", e)
	return &ucenter.MemberWallet{}, e
}

func e1() error {
	return errors.New("aaaaaaa")
}
