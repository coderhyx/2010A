package svc

import (
	"fmt"
	"goods_srv/internal/config"
	"goods_srv/internal/dao"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type ServiceContext struct {
	Config   config.Config
	GoodsDao *dao.GoodsDao
}

func NewServiceContext(c config.Config) *ServiceContext {

	//db := db.GormConn(c.Mysql.DataSource)
	dao.InitMySQL(c.Mysql.DataSource)
	db := dao.GetDB()
	nacosInit()
	return &ServiceContext{
		Config:   c,
		GoodsDao: dao.NewGoodsDao(db),
	}
}

func nacosInit() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("192.168.43.44", 8848, constant.WithContextPath("/nacos")),
	}
	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("c255669f-545d-4978-bdfa-72cd5f34d5be"),
		//constant.WithTimeoutMs(5000),
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)
	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	//_, err = client.PublishConfig(vo.ConfigParam{
	//	DataId:  "test-data",
	//	Group:   "test-group",
	//	Content: "hello world!",
	//})
	//content, err := client.GetConfig(vo.ConfigParam{
	//	DataId: "test-data",
	//	Group:  "test-group",
	//})
	//fmt.Println("GetConfig,config :" + content)

	content1, err := client.GetConfig(vo.ConfigParam{
		DataId: "goods_web.json",
		Group:  "dev",
	})
	fmt.Println("GetConfig,config :" + content1)
}

func init2() {
	sc := []constant.ServerConfig{{
		IpAddr: "127.0.0.1",
		Port:   8848,
	}}

	cc := constant.ClientConfig{
		NamespaceId:         "c255669f-545d-4978-bdfa-72cd5f34d5be", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "debug",
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	content1, err := client.GetConfig(vo.ConfigParam{
		DataId: "goods_web.json",
		Group:  "dev",
	})
	fmt.Println("GetConfig,config :" + content1)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(content) //字符串 - yaml
	//err = configClient.ListenConfig(vo.ConfigParam{
	//	DataId: "user-dev",
	//	Group:  "dev",
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("配置文件发生了变化...")
	//		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	//	},
	//})
}
