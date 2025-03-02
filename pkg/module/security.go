package module

import (
	"github.com/onerciller/fullstack-golang-template/pkg/config"
	"github.com/onerciller/fullstack-golang-template/pkg/jwt"
	"go.uber.org/fx"
)

type SecurityModule struct {
	fx.In
}

func (m *SecurityModule) Provide() fx.Option {
	return fx.Module("security",
		fx.Provide(func(cfg config.ConfigProvider) jwt.Jwt {
			return jwt.NewJWTService(cfg.GetString("jwt.secret_key"))
		}),
	)
}
