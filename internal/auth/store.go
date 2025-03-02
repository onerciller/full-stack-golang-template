package auth

import (
	"context"
	stderrors "errors"

	"github.com/onerciller/fullstack-golang-template/internal/shared"
	"gorm.io/gorm"
)

// Store implements UserStore interface
type Store struct {
	db *gorm.DB
}

// New creates a new instance of UserStore
func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// Implementation of UserStore interface methods
func (s *Store) FindByID(ctx context.Context, id uint) (*UserEntity, error) {
	var user UserEntity
	if err := s.db.First(&user, id).Error; err != nil {
		if stderrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrUserNotFound.ToNotFoundAppError()
		}
		return nil, err
	}
	return &user, nil
}

func (s *Store) FindByEmail(ctx context.Context, email string) (*UserEntity, error) {
	var user UserEntity
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if stderrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrUserNotFound.ToNotFoundAppError()
		}
		return nil, err
	}
	return &user, nil
}

func (s *Store) FindByUsername(ctx context.Context, username string) (*UserEntity, error) {
	var user UserEntity
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, shared.ErrUserNotFound.ToNotFoundAppError()
	}
	return &user, nil
}

func (s *Store) FindAll(ctx context.Context) ([]*UserEntity, error) {
	var users []*UserEntity
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Store) Create(ctx context.Context, user *UserEntity) error {
	return s.db.Create(user).Error
}

func (s *Store) Update(ctx context.Context, user *UserEntity) error {
	if err := s.db.Save(user).Error; err != nil {
		if stderrors.Is(err, gorm.ErrRecordNotFound) {
			return shared.ErrUserNotFound.ToNotFoundAppError()
		}
		return err
	}
	return nil
}

func (s *Store) Delete(ctx context.Context, id uint) error {
	result := s.db.Delete(&UserEntity{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return shared.ErrUserNotFound.ToNotFoundAppError()
	}
	return nil
}
