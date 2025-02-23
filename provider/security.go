package provider

import (
	"github.com/onerciller/fullstack-golang-template/pkg/config"
	"github.com/onerciller/fullstack-golang-template/pkg/jwt"
)

type Security struct{}

func (c *Security) Provide(cfg config.ConfigProvider) jwt.Jwt {
	return jwt.NewJWTService(cfg.GetString("jwt.secret_key"))
}
