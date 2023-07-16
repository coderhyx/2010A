package svc

import (
	"ucenter_srv/internal/config"
	"ucenter_srv/internal/model"

	"github.com/zeromicro/go-zero/core/stores/cache"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	MemberModel model.MemberModel
	Cache       cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("wtcoin"),
		nil,
		func(o *cache.Options) {})
	return &ServiceContext{
		Config:      c,
		MemberModel: model.NewMemberModel(conn),
		Cache:       redisCache,
	}
}
