package controllers

import (
	"configer-service/internal/core"
	"configer-service/internal/models"
	"configer-service/internal/response"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	user core.UserServiceI
}

func NewUserHandler(userI core.UserServiceI) *UserHandler {
	return &UserHandler{
		user: userI,
	}
}

func (h UserHandler) CreateUser(e echo.Context) error {
	user := &models.User{}
	if err := h.bindUser(e, user); err != nil {
		return err
	}
	if err := h.user.Create(user); err != nil {
		return response.Error(e, err)
	}
	return response.OK(e, http.StatusCreated, user)
}

func (h UserHandler) GetUsersList(e echo.Context) error {
	users, err := h.user.List()
	if err != nil {
		return response.Error(e, err)
	}
	return response.OK(e, http.StatusOK, users)
}

func (h UserHandler) GetUserByName(e echo.Context) error {
	user := &models.User{}
	if err := h.bindUser(e, user); err != nil {
		return err
	}
	if err := h.user.Get(user); err != nil {
		return response.Error(e, err)
	}
	return response.OK(e, http.StatusOK, user)
}

func (h UserHandler) bindUser(e echo.Context, user *models.User) error {
	if user == nil {
		log.Print("Warning: pointer to struct expected but nil got")
		user = &models.User{}
	}
	if err := e.Bind(user); err != nil {
		return response.Error(e, err)
	}
	return nil
}
