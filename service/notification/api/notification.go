// package main

// import (
// 	"flag"
// 	"fmt"

// 	"roomrover/service/notification/api/internal/config"
// 	"roomrover/service/notification/api/internal/handler"
// 	"roomrover/service/notification/api/internal/svc"

// 	"github.com/zeromicro/go-zero/core/conf"
// 	"github.com/zeromicro/go-zero/rest"
// )

// var configFile = flag.String("f", "etc/notification-api.yaml", "the config file")

// func main() {
// 	flag.Parse()

// 	var c config.Config
// 	conf.MustLoad(*configFile, &c)

// 	server := rest.MustNewServer(c.RestConf)
// 	defer server.Stop()

// 	ctx := svc.NewServiceContext(c)
// 	handler.RegisterHandlers(server, ctx)

// 	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
// 	server.Start()
// }

package api

import (
	"flag"
	"roomrover/service/notification/api/internal/config"
	"roomrover/service/notification/api/internal/handler"
	"roomrover/service/notification/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/notification-api.yaml", "the config file")

type NotificationService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewNotificationService(server *rest.Server) *NotificationService {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return &NotificationService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (s *NotificationService) Start() error {
	return nil
}
