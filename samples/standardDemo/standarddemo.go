package main

import (
	"flag"
	"fmt"

	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/config"
	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/handler"
	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/standarddemo-api.yaml", "the config file")

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
