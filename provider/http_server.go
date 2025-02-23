package provider

import (
	"github.com/onerciller/fullstack-golang-template/pkg/config"
	"github.com/onerciller/fullstack-golang-template/pkg/httpserver"
	"go.uber.org/fx"
)

type HttpServer struct{}

func (c *HttpServer) Provide(lc fx.Lifecycle, cfg config.ConfigProvider) *httpserver.HttpServer {
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
}
