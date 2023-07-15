package svc

import (
	"ucenter_api/internal/config"
	"ucenter_api/internal/middleware"
	"ucenter_srv/ucenterclient"

	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	AuthInterceptor rest.Middleware
	UCenterSrv      ucenterclient.Ucenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		AuthInterceptor: middleware.NewAuthInterceptorMiddleware().Handle,
		UCenterSrv:      ucenterclient.NewUcenter(zrpc.MustNewClient(c.UCenterSrv)),
	}
}
