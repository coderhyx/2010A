package svc

import (
	"2010A/common/db"
	"goods_srv/internal/config"
	"goods_srv/internal/dao"
)

type ServiceContext struct {
	Config   config.Config
	GoodsDao *dao.GoodsDao
	Es       *db.ElasticsearchClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	//db := db.GormConn(c.Mysql.DataSource)
	dao.InitMySQL(c.Mysql.DataSource)
	mysqlDao := dao.GetDB()
	es, _ := db.NewElasticsearchClient(c.EsDB.DataSource)
	return &ServiceContext{
		Config:   c,
		GoodsDao: dao.NewGoodsDao(mysqlDao),
		Es:       es,
	}
}
