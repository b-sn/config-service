package controller

import (
	"configer-service/internal/core"
	"configer-service/internal/custom"
	"configer-service/internal/response"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	ConfigHandler struct {
		cfg core.ConfigServiceI
	}
	ConfigRequest struct {
		AppName string `param:"app_name" validate:"required"`
		EnvName string `param:"env_name" validate:"required"`
		UserID  string `json:"-"`
	}
)

func NewConfigHandler(cfgI core.ConfigServiceI) *ConfigHandler {
	return &ConfigHandler{
		cfg: cfgI,
	}
}

func (c ConfigHandler) CreateConfig(ctx echo.Context) error {

	cfgReq, err := c.bindConfigParams(ctx)
	if err != nil {
		return ctx.JSON(response.Error(err))
	}

	return ctx.JSON(response.OKApp(http.StatusOK, app))

}

func (c ConfigHandler) bindConfigParams(ctx echo.Context) (ConfigRequest, error) {

	userID, ok := ctx.Get("user_id").(string)
	if !ok || userID == "" {
		return ConfigRequest{}, custom.NewRequestError(
			errors.New("echo context contains invalid user id"),
		)
	}

	var cfg ConfigRequest

	if err := ctx.Bind(&cfg); err != nil {
		return ConfigRequest{}, custom.NewRequestError(err)
	}

	if err := ctx.Validate(&cfg); err != nil {
		return ConfigRequest{}, custom.NewRequestError(err)
	}

	cfg.AppName = strings.TrimRight(cfg.AppName, "/")

	// Check if application belongs to the user

	cfg.EnvName = strings.TrimRight(cfg.EnvName, "/")
	cfg.UserID = userID

	return cfg, nil
}
