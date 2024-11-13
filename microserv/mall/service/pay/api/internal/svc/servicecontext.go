package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"plan_go/microserv/mall/service/pay/api/internal/config"
	"plan_go/microserv/mall/service/pay/rpc/payclient"
)

type ServiceContext struct {
	Config config.Config

	PayRpc payclient.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}