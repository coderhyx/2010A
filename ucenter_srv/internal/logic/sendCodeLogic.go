package logic

import (
	"2010A/common/feature"
	"2010A/common/utils"
	"context"
	"errors"
	"time"
	"ucenter_srv/internal/model"

	"github.com/zeromicro/go-zero/core/logc"

	"ucenter_srv/internal/svc"
	"ucenter_srv/pb/ucenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *ucenter.CodeReq) (*ucenter.NoRes, error) {
	//1.生成验证码
	code := utils.Rand4Num()
	//2.短信平台发送验证
	sms := feature.UniSMS{
		AccessKeyId:  l.svcCtx.Config.SMS.AccessKeyId,
		AccessSecret: l.svcCtx.Config.SMS.AccessSecret,
		Signature:    l.svcCtx.Config.SMS.Signature,
		TemplateId:   l.svcCtx.Config.SMS.TemplateId,
	}
	err := sms.SendCode(in.Phone, code)
	if err != nil {
		return nil, errors.New("网络原因导致验证码发送失败")
	}
	//2.将验证码存入redis，过期时间5分钟
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	l.svcCtx.Cache.SetWithExpireCtx(ctx, model.SendCodeCacheKey+in.Phone, code, 5*time.Minute)
	//4.返回成功
	if err != nil {
		logc.Errorf(l.ctx, "send code fail, because %s", err)
		return nil, errors.New("缓存原因，验证码发送失败")
	}
	logc.Info(l.ctx, "send code success")
	return &ucenter.NoRes{}, nil
}
