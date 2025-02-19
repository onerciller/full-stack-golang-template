package database

import (
	"github.com/onerciller/fullstack-golang-template/pkg/config"
	"github.com/onerciller/fullstack-golang-template/pkg/database/postgres"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvidePostgres(i *do.Injector) (*gorm.DB, error) {
	cfg := do.MustInvoke[config.ConfigProvider](i)

	db := postgres.New(
		postgres.WithHost(cfg.GetString("db.host")),
		postgres.WithPort(cfg.GetString("db.port")),
		postgres.WithUser(cfg.GetString("db.user")),
		postgres.WithPass(cfg.GetString("db.password")),
		postgres.WithDbName(cfg.GetString("db.dbname")),
	)
	return db, nil
}
