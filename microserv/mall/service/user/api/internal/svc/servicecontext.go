package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"plan_go/microserv/mall/service/user/api/internal/config"
	"plan_go/microserv/mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
