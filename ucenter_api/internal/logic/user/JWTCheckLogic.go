package user

import (
	"2010A/common/feature"
	"context"
	"fmt"

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
	r, err := feature.ParseToken(req.Token, l.svcCtx.Config.JWT.AccessSecret)
	fmt.Println("--------->", r)
	if err != nil {
		logx.Error(err)
		return nil, nil
	}
	return &types.BoolResponse{
		Res: true,
	}, nil
}
