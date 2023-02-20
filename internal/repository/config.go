package repository

import (
	"configer-service/internal/model"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConfigRepo struct {
	db *gorm.DB
}

type ConfigI interface {
	Create(cfg *model.Config) error
}

func NewConfigRepo(db *gorm.DB) *ConfigRepo {
	db.AutoMigrate((new(model.Config)))
	return &ConfigRepo{
		db: db,
	}
}

func (c ConfigRepo) Create(cfg *model.Config) error {
	cfg.ID = uuid.NewString()
	if err := c.db.Create(cfg).Error; err != nil {
		return fmt.Errorf("cannot add config to DB: %v", err)
	}
	return nil
}
