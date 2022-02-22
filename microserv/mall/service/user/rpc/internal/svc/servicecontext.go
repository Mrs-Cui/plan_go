package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	model2 "plan_go/plan_go/microserv/mall/service/user/model"
	config2 "plan_go/plan_go/microserv/mall/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config2.Config

	UserModel model2.UserModel
}

func NewServiceContext(c config2.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model2.NewUserModel(conn, c.CacheRedis),
	}
}
