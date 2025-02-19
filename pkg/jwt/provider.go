package jwt

import (
	"github.com/onerciller/fullstack-golang-template/pkg/config"
	"github.com/samber/do"
)

func ProvideJWTService(i *do.Injector) (Jwt, error) {
	cfg := do.MustInvoke[config.ConfigProvider](i)
	return NewJWTService(cfg.GetString("jwt.secret_key")), nil
}
