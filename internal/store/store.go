package store

import (
	"context"

	"github.com/onerciller/fullstack-golang-template/internal/entity"
	"github.com/onerciller/fullstack-golang-template/internal/store/user"
	"gorm.io/gorm"
)

// UserRepository defines the interface for user-related database operations
type UserStore interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, id uint) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uint) error
}

// ProvideUserStore provides a user store
func NewUserStore(db *gorm.DB) UserStore {
	return user.New(db)
}
