package module

import (
	"github.com/onerciller/fullstack-golang-template/pkg/config"
	"github.com/onerciller/fullstack-golang-template/pkg/httpserver"
	"go.uber.org/fx"
)

type HttpServerModule struct {
	fx.In
}

func (m *HttpServerModule) Provide() fx.Option {
	return fx.Module("httpserver",
		fx.Provide(func(cfg config.ConfigProvider) *httpserver.HttpServer {
			server := httpserver.New(
				httpserver.WithServerHeader(cfg.GetString("server.header")),
				httpserver.WithAppName(cfg.GetString("server.name")),
				httpserver.WithPort(cfg.GetString("server.port")),
				httpserver.WithMiddlewares(
					httpserver.WithCORS(),
					httpserver.WithHealthCheck("/health"),
				),
			)
			return server
		}),
	)
}
