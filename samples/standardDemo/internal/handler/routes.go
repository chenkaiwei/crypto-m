// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Multi1Request, serverCtx.Multi1Response},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/multiDemo1",
					Handler: MultiDemo1Handler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Multi2Request},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/multiDemo2",
					Handler: MultiDemo2Handler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Muti1Manual},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/manualDemo",
					Handler: ManualDemoHandler(serverCtx),
				},
			}...,
		),
	)
}