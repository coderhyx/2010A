package svc

import (
	"ucenter_srv/internal/config"
	"ucenter_srv/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	MemberModel model.MemberModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		MemberModel: model.NewMemberModel(conn, c.CacheRedis),
	}
}
