package repository

import (
	"configer-service/internal/model"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

type UserI interface {
	Create(user *model.User) error
	Get(user *model.User) error
	Exists(userName string) (bool, error)
	ExistsAndActive(userID string) (bool, error)
	Find(user *model.User) ([]*model.User, error)
	// Deactivate() error
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	db.AutoMigrate(new(model.User))
	return &UserRepo{
		db: db,
	}
}

func (r UserRepo) Create(user *model.User) error {
	user.Name = strings.ToLower(user.Name)
	user.ID = uuid.NewString()
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("cannot save user: %w", err)
	}
	return nil
}

func (r UserRepo) Find(user *model.User) ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Where(user).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepo) Get(user *model.User) error {
	if err := r.db.Where(user).Take(user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("error while select user: %w", err)
		}
	}
	return nil
}

func (r UserRepo) ExistsAndActive(userID string) (bool, error) {
	user := model.User{
		ID:       userID,
		IsActive: true,
	}
	if err := r.db.Where(&user).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("error while select user: %w", err)
	}
	return true, nil
}

func (r UserRepo) Exists(name string) (bool, error) {
	user := model.User{
		Name: name,
	}
	if err := r.db.Where(&user).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("error while select user: %w", err)
	}
	return true, nil
}
