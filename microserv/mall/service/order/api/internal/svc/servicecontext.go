package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"plan_go/microserv/mall/service/order/api/internal/config"
	"plan_go/microserv/mall/service/order/rpc/orderclient"
	"plan_go/microserv/mall/service/product/rpc/productclient"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc orderclient.Order
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
