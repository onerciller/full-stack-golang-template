package provider

import (
	"github.com/onerciller/fullstack-golang-template/pkg/config"
	"github.com/onerciller/fullstack-golang-template/pkg/database/postgres"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Database struct{}

func (c *Database) Provide(lc fx.Lifecycle, cfg config.ConfigProvider) *gorm.DB {

	posgres := postgres.New(
		postgres.WithHost(cfg.GetString("db.host")),
		postgres.WithPort(cfg.GetString("db.port")),
		postgres.WithUser(cfg.GetString("db.user")),
		postgres.WithPass(cfg.GetString("db.password")),
		postgres.WithDbName(cfg.GetString("db.dbname")),
	)

	return posgres
}
