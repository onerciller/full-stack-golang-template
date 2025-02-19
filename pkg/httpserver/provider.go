package httpserver

import (
	"github.com/onerciller/fullstack-golang-template/pkg/config"
	"github.com/samber/do"
)

func Provide(di *do.Injector) (*HttpServer, error) {
	config := do.MustInvoke[config.ConfigProvider](di)
	httpServer := New(
		WithServerHeader(config.GetString("app.name")),
		WithAppName(config.GetString("app.name")),
		WithPort(config.GetString("server.port")),
		WithMiddlewares(
			WithCORS(),
			WithHealthCheck("/health"),
		),
	)
	return httpServer, nil
}
