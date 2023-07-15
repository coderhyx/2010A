package user

import (
	"context"

	"ucenter_api/internal/svc"
	"ucenter_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JWTCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJWTCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JWTCheckLogic {
	return &JWTCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JWTCheckLogic) JWTCheck(req *types.JWTRequest) (resp *types.BoolResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
