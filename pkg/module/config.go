package module

import (
	"github.com/onerciller/fullstack-golang-template/pkg/config"
	"go.uber.org/fx"
)

type ConfigModule struct {
	fx.In
}

func (m *ConfigModule) Provide() fx.Option {
	return fx.Module("config",
		fx.Provide(func() config.ConfigProvider {
			return config.New(
				config.WithPath("./configs"),
				config.WithConfigName("config"),
				config.WithConfigType("yaml"),
			)
		}),
	)
}
