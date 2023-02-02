package core

import (
	"configer-service/internal/custom"
	"configer-service/internal/models"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
)

// const hiddenAppTokenReplace = "*hidden*"

type AppServiceI interface {
	Create(*models.App) error
	Get(*models.App) error
	// List() ([]*models.User, error)
}

type appService struct {
	repo models.AppRepository
}

func NewAppService(r models.AppRepository) appService {
	return appService{
		repo: r,
	}
}

func (a appService) Create(app *models.App) error {
	if app.Name == "" {
		return custom.RequestError{
			Text: fmt.Sprintf("name expected to create application, got: %#v", app),
		}
	}
	if err := a.repo.Get(app); err != nil {
		log.Errorf("getting application error: %w, app: %#v", err, app)
		return errors.New("unexpected get application error")
	}
	if app.ID != 0 {
		return custom.DuplicateError{
			Text: fmt.Sprintf("Application with name [%s] already exists, try another name", app.Name),
		}
	}
	if err := a.repo.Create(app); err != nil {
		log.Errorf("creating application error: %w, app: %#v", err, app)
		return errors.New("unexpected create application error")
	}
	return nil
}

func (a appService) Get(app *models.App) error {
	if app.Name == "" && app.ID == 0 && app.Token == "" {
		return custom.RequestError{
			Text: fmt.Sprintf("parameters to get application expected, got: %#v", app),
		}
	}
	if err := a.repo.Get(app); err != nil {
		log.Errorf("getting application error: %w, app: %#v", err, app)
		return errors.New("unexpected get application error, see logs")
	}
	return nil
}
