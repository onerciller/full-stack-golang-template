package provider

import (
	"go.uber.org/fx"
)

func Bootstrap() *fx.App {
	appConfig := &AppConfig{}
	security := &Security{}
	database := &Database{}
	store := &Store{}
	httpServer := &HttpServer{}
	lifecycle := &Lifecycle{}
	api := &Api{}
	return fx.New(
		fx.Provide(appConfig.Provide),
		fx.Provide(httpServer.Provide),
		fx.Provide(database.Provide),
		fx.Provide(store.Provide),
		fx.Provide(security.Provide),
		fx.Invoke(api.Provide),
		fx.Invoke(lifecycle.Provide),
	)
}
