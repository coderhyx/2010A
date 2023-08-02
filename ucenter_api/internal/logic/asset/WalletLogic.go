package asset

import (
	"2010A/common/tools"
	"context"
	"ucenter_api/internal/svc"
	"ucenter_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletLogic {
	return &WalletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WalletLogic) Wallet(req *types.WalletRequest) (resp *types.WalletResponse, err error) {
	logx.Info("get my wallet success ........")
	tools.FormatTime(334382943298492)
	logx.Info("调用通用模块的时间格式化函数")
	return
}
