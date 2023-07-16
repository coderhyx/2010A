package logic

import (
	"context"
	//"errors"

	"github.com/pkg/errors"

	"ucenter_srv/internal/svc"
	"ucenter_srv/pb/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type ErrorDemoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewErrorDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ErrorDemoLogic {
	return &ErrorDemoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// error demo
func (l *ErrorDemoLogic) ErrorDemo(in *ucenter.ErrorDemoReq) (*ucenter.ErrorDemoRes, error) {
	if r := l.demo1(); r != nil {
		logx.Errorf("捕获error: %+v", r)
		//logx.Errorf("在最外层接收到了err: %s", r)
	}
	//r := l.demo1()
	//logx.Errorf("在最外层接收到了err: %v", r)
	//logx.Error("aaaaaaaaaaaaaaaaaaaaaaaa")
	return &ucenter.ErrorDemoRes{}, nil
}

func (l *ErrorDemoLogic) demo1() error {
	return errors.Wrap(l.demo2(), "demo1")
}

func (l *ErrorDemoLogic) demo2() error {
	return errors.Wrap(l.demo3(), "demo2")
}

func (l *ErrorDemoLogic) demo3() error {

	//return errors.New("躲在第三层，发现不了吧")
	return errors.Wrap(errors.New("躲在第三层，发现不了吧"), "demo3")
}
