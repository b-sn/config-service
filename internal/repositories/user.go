package repositories

import (
	"configer-service/internal/models"
	"configer-service/pkg/utils"
	"encoding/base64"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r UserRepo) Create(user *models.User) error {
	user.IsActive = true
	user.Token = base64.StdEncoding.EncodeToString(utils.GenerateRandData(64))
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("cannot save user: %v", err)
	}
	return nil
}

func (r UserRepo) Find(user *models.User) ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Where(user).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepo) Get(user *models.User) error {
	if err := r.db.Where(user).Take(user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("error while select user: %v", err)
		}
	}
	return nil
}
