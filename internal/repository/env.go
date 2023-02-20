package repository

import (
	"configer-service/internal/model"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EnvRepo struct {
	db *gorm.DB
}

type EnvI interface {
	Create(env *model.Env) error
}

func NewEnvRepo(db *gorm.DB) *EnvRepo {
	db.AutoMigrate(new(model.Env))
	return &EnvRepo{
		db: db,
	}
}

func (e EnvRepo) Create(env *model.Env) error {
	env.ID = uuid.NewString()
	if err := e.db.Create(env).Error; err != nil {
		return fmt.Errorf("cannot add environment to DB: %v", err)
	}
	return nil
}
