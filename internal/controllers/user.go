package controllers

import (
	"configer-service/internal/models"
	"configer-service/internal/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userRepo models.UserRepository
}

const (
	paramsError     = 1001
	createUserError = 1002
)

func NewUserHandler(userRepo models.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (h *UserHandler) AddUser(e echo.Context) error {
	var user models.User
	err := e.Bind(&user)
	if err != nil {
		return e.JSON(
			http.StatusBadRequest,
			response.ReturnErrorJSON(paramsError, "Param user_name error", err))
	}
	err = h.userRepo.Create(&user)
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			response.ReturnErrorJSON(createUserError, "User creation error", err))
	}
	return e.JSON(response.ReturnOkJSON(user))
}

func (h *UserHandler) UserList(e echo.Context) error {
	users, err := h.userRepo.List()
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			response.ReturnErrorJSON(createUserError, "Getting user list error", err))
	}
	return e.JSON(response.ReturnOkJSON(users))
}
