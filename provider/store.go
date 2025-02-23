package provider

import (
	"github.com/onerciller/fullstack-golang-template/internal/store"
	"gorm.io/gorm"
)

type Store struct{}

func (c *Store) Provide(db *gorm.DB) store.UserStore {
	return store.NewUserStore(db)
}
