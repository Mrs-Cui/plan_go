package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"plan_go/microserv/mall/service/product/api/internal/config"
	"plan_go/microserv/mall/service/product/rpc/productclient"
)

type ServiceContext struct {
	Config config.Config

	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
