package main

import (
	"flag"
	"fmt"

	"plan_go/microserv/mall/service/order/api/internal/config"
	"plan_go/microserv/mall/service/order/api/internal/handler"
	"plan_go/microserv/mall/service/order/api/internal/svc"

	_ "github.com/dtm-labs/driver-gozero"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
