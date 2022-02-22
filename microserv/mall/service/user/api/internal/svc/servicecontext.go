package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	config2 "plan_go/plan_go/microserv/mall/service/user/api/internal/config"
	userclient2 "plan_go/plan_go/microserv/mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config2.Config

	UserRpc userclient2.User
}

func NewServiceContext(c config2.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient2.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
