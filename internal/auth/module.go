package auth

import (
	"github.com/onerciller/fullstack-golang-template/pkg/httpserver"
	"go.uber.org/fx"
)

type Module struct {
	fx.In
	Store   *Store
	Handler *Handler
}

func (m *Module) Provide() fx.Option {
	return fx.Module("auth",
		fx.Provide(NewStore),
		fx.Provide(func(store *Store) UserStore {
			return store
		}),
		fx.Provide(NewHandler),
	)
}

func (m *Module) Invoke() fx.Option {
	return fx.Invoke(func(handler *Handler, httpServer *httpserver.HttpServer) {
		handler.RegisterRoutes(httpServer.FiberApp)
	})
}
