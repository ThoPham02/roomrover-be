package main

import (
	"flag"
	"fmt"
	"roomrover/config"

	middleware "github.com/muhfajar/go-zero-cors-middleware"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	account "roomrover/service/account/api"
)

var configFile = flag.String("f", "etc/server.yaml", "the config file")

// @BasePath  
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	cors := middleware.NewCORSMiddleware(&middleware.Options{})

	server := rest.MustNewServer(c.RestConf, rest.WithNotAllowedHandler(cors.Handler()))
	server.Use(cors.Handle)
	defer server.Stop()

	accountService := account.NewAccountService(server)
	accountService.Start()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	server.Start()
}
