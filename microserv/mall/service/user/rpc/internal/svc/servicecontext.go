package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"plan_go/microserv/mall/service/user/model"
	"plan_go/microserv/mall/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
