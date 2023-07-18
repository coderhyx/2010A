package svc

import (
	"goods_srv/internal/config"
	"goods_srv/internal/dao"
)

type ServiceContext struct {
	Config   config.Config
	GoodsDao *dao.GoodsDao
}

func NewServiceContext(c config.Config) *ServiceContext {

	//db := db.GormConn(c.Mysql.DataSource)
	dao.InitMySQL(c.Mysql.DataSource)

	return &ServiceContext{
		Config:   c,
		GoodsDao: dao.NewGoodsDao(),
	}
}
