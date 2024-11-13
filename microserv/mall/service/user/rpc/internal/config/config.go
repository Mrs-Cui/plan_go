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
	Salt string
	Auth          bool
	StrictControl bool
	Redis         struct {
		Key  string
		Host string
		Type string
		Pass string
	}
	Prometheus struct{
		Host string
		Port int64
		Path string
	}
}
