package repositories

import (
	"configer-service/internal/models"
	"crypto/rand"
	"encoding/base64"
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

func (r *UserRepo) Create(user *models.User) error {
	user.IsActive = true

	b := make([]byte, 64)
	if _, err := rand.Read(b); err != nil {
		return fmt.Errorf("cannot generate token: %v", err)
	}
	user.Token = base64.StdEncoding.EncodeToString(b)

	result := r.db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("cannot save user: %v", result.Error)
	}

	return nil
}

func (r *UserRepo) List() ([]*models.User, error) {
	var users []*models.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return []*models.User{}, fmt.Errorf("cannot select user list: %v", result.Error)
	}
	for _, user := range users {
		user.Token = "*hidden*"
	}
	return users, nil
}
