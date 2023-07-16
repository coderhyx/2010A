package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	SMS        UniCodeConfig
}

type UniCodeConfig struct {
	AccessKeyId  string
	AccessSecret string
	Signature    string
	TemplateId   string
}
