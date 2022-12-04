package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/chenkaiwei/crypto-m/samples/simpleDemo/internal/config"
	"github.com/chenkaiwei/crypto-m/samples/simpleDemo/internal/handler"
	"github.com/chenkaiwei/crypto-m/samples/simpleDemo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/simpledemo-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		//logx.Info("ErrorHandler--", err)

		logx.Errorf("print errStack--%+v\n", err) //打印 包装时的调用栈

		return http.StatusInternalServerError, err

	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
