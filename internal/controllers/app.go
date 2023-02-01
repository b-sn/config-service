package controllers

import (
	"configer-service/internal/core"
	"configer-service/internal/models"
	"configer-service/internal/response"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppHandler struct {
	app core.AppServiceI
}

func NewAppHandler(appI core.AppServiceI) *AppHandler {
	return &AppHandler{
		app: appI,
	}
}

func (h AppHandler) CreateApp(e echo.Context) error {
	app := &models.App{}
	if err := h.bindApp(e, app); err != nil {
		return err
	}

	e.Path()

	if err := h.app.Create(app); err != nil {
		return response.Error(e, err)
	}
	return response.OK(e, http.StatusCreated, app)
}

func (h AppHandler) bindApp(e echo.Context, app *models.App) error {
	if app == nil {
		log.Print("Warning: pointer to struct expected but nil got")
		app = &models.App{}
	}
	if err := e.Bind(app); err != nil {
		return response.Error(e, err)
	}

	// TODO:

	return nil
}
