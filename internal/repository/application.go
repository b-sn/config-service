package repository

import (
	"configer-service/internal/model"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppRepo struct {
	db *gorm.DB
}

type ApplicationI interface {
	Create(app *model.App) error
	Get(app *model.App) error
	// FindByName(name string) (*App, error)
	// FindByToken(token string) (*App, error)
	// Update(params map[string]string) error
	// Deactivate() error
}

func NewAppRepo(db *gorm.DB) *AppRepo {
	db.AutoMigrate(new(model.App))
	return &AppRepo{
		db: db,
	}
}

func (a AppRepo) Create(app *model.App) error {
	app.Name = strings.ToLower(app.Name)
	app.ID = uuid.NewString()
	// app.Token = base64.StdEncoding.EncodeToString(utils.GenerateRandData(64))
	if err := a.db.Create(app).Error; err != nil {
		return fmt.Errorf("cannot add application to DB: %v", err)
	}
	return nil
}

func (a AppRepo) Get(app *model.App) error {
	if err := a.db.Where(app).Take(app).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("error while select application: %v", err)
		}
	}
	return nil
}
