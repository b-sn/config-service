package repositories

import (
	"configer-service/internal/models"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/labstack/gommon/log"
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
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	for _, user := range users {
		user.Token = "*hidden*"
	}
	return users, nil
}

func (r *UserRepo) GetUserByName(name string) ([]*models.User, error) {
	var users []*models.User
	if result := r.db.Where("`user_name` = ?", name).Find(&users); result.Error != nil {
		log.Errorf("cannot select user list: %v", result.Error)
		return nil, fmt.Errorf("cannot select user list")
	}
	log.Panic(users)

	return users, nil
}
