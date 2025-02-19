package user

import (
	"context"
	stderrors "errors"

	"github.com/onerciller/fullstack-golang-template/internal/entity"
	"github.com/onerciller/fullstack-golang-template/internal/errors"
	"gorm.io/gorm"
)

// Store implements UserStore interface
type Store struct {
	db *gorm.DB
}

// New creates a new instance of UserStore
func New(db *gorm.DB) *Store {
	return &Store{db: db}
}

// Implementation of UserStore interface methods
func (s *Store) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	if err := s.db.First(&user, id).Error; err != nil {
		if stderrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrUserNotFound.ToNotFoundAppError()
		}
		return nil, err
	}
	return &user, nil
}

func (s *Store) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if stderrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrUserNotFound.ToNotFoundAppError()
		}
		return nil, err
	}
	return &user, nil
}

func (s *Store) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.ErrUserNotFound.ToNotFoundAppError()
	}
	return &user, nil
}

func (s *Store) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Store) Create(ctx context.Context, user *entity.User) error {
	return s.db.Create(user).Error
}

func (s *Store) Update(ctx context.Context, user *entity.User) error {
	if err := s.db.Save(user).Error; err != nil {
		if stderrors.Is(err, gorm.ErrRecordNotFound) {
			return errors.ErrUserNotFound.ToNotFoundAppError()
		}
		return err
	}
	return nil
}

func (s *Store) Delete(ctx context.Context, id uint) error {
	result := s.db.Delete(&entity.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.ErrUserNotFound.ToNotFoundAppError()
	}
	return nil
}

func (s *Store) UpdateRefreshToken(ctx context.Context, userID uint, token string) error {
	return s.db.Model(&entity.User{}).Where("id = ?", userID).Update("refresh_token", token).Error
}
