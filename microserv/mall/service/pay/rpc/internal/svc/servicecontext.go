package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"plan_go/microserv/mall/service/order/rpc/orderclient"
	"plan_go/microserv/mall/service/pay/model"
	"plan_go/microserv/mall/service/pay/rpc/internal/config"
	"plan_go/microserv/mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	PayModel model.PayModel

	UserRpc  userclient.User
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}