package svc

import (
	common "2010A/common/ctl"
	"2010A/common/db"
	"fmt"
	"ucenter_srv/internal/config"
	"ucenter_srv/internal/model"

	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/zeromicro/go-zero/core/stores/cache"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	MemberModel model.MemberModel
	Cache       cache.Cache
	ES          db.ElasticsearchClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("wtcoin"),
		nil,
		func(o *cache.Options) {})
	es, _ := db.NewElasticsearchClient()
	initNacos()
	return &ServiceContext{
		Config:      c,
		MemberModel: model.NewMemberModel(conn),
		Cache:       redisCache,
		ES:          *es,
	}
}

func initNacos() {
	//从nacos中读取配置信息
	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "c255669f-545d-4978-bdfa-72cd5f34d5be", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}
	nc, _ := common.NewNacosClient(sc, cc)
	str, _ := nc.GetConfig("goods_srv.json", "dev")
	fmt.Println("--------------->main:", str)
}
