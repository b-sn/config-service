package controller

import (
	"configer-service/internal/core"
	"configer-service/internal/custom"
	"configer-service/internal/model"
	"configer-service/internal/response"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	AppHandler struct {
		app core.AppServiceI
	}
	AppRequest struct {
		Name         string `param:"app_name" json:"app_name" validate:"required,max=100"`
		UserID       string `json:"-"`
		IsInactive   bool   `json:"is_inactive,omitempty"`
		IsVersioning bool   `json:"is_versions,omitempty"`
	}
)

func NewAppHandler(appI core.AppServiceI) *AppHandler {
	return &AppHandler{
		app: appI,
	}
}

func (h AppHandler) CreateApp(ctx echo.Context) error {

	appReq, err := h.bindAppParams(ctx)
	if err != nil {
		return ctx.JSON(response.Error(err))
	}

	app := &model.App{
		UserID:     appReq.UserID,
		Name:       appReq.Name,
		IsActive:   !appReq.IsInactive,
		IsVersions: appReq.IsVersioning,
	}

	if err := h.app.Create(app); err != nil {
		return ctx.JSON(response.Error(err))
	}
	return ctx.JSON(response.OKApp(http.StatusOK, app))
}

func (h AppHandler) GetAppByName(ctx echo.Context) error {
	appReq, err := h.bindAppParams(ctx)
	if err != nil {
		return ctx.JSON(response.Error(err))
	}

	app := &model.App{
		Name: appReq.Name,
	}
	if err := h.app.Get(app); err != nil {
		return ctx.JSON(response.Error(err))
	}

	return ctx.JSON(response.OKApp(http.StatusOK, app))
}

func (h AppHandler) bindAppParams(ctx echo.Context) (AppRequest, error) {

	userID, ok := ctx.Get("user_id").(string)
	if !ok || userID == "" {
		return AppRequest{}, custom.NewRequestError(
			errors.New("echo context contains invalid user id"),
		)
	}

	var app AppRequest

	if err := ctx.Bind(&app); err != nil {
		return AppRequest{}, custom.NewRequestError(err)
	}

	if err := ctx.Validate(&app); err != nil {
		return AppRequest{}, custom.NewRequestError(err)
	}

	app.Name = strings.TrimRight(app.Name, "/")
	app.UserID = userID

	return app, nil
}
