package repositories

import (
	"configer-service/internal/models"
	"configer-service/pkg/utils"
	"encoding/base64"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AppRepo struct {
	db *gorm.DB
}

func NewAppRepo(db *gorm.DB) *AppRepo {
	return &AppRepo{
		db: db,
	}
}

func (a AppRepo) Create(app *models.App) error {
	app.IsActive = true
	app.Token = base64.StdEncoding.EncodeToString(utils.GenerateRandData(64))
	if err := a.db.Create(app).Error; err != nil {
		return fmt.Errorf("cannot save app: %v", err)
	}
	return nil
}

func (a AppRepo) Get(app *models.App) error {
	if err := a.db.Where(app).Take(app).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("error while select application: %v", err)
		}
	}
	return nil
}
