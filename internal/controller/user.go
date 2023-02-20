package controller

import (
	"configer-service/internal/core"
	"configer-service/internal/custom"
	"configer-service/internal/model"
	"configer-service/internal/response"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	UserHandler struct {
		user core.UserServiceI
	}
	UserRequest struct {
		Name       string `param:"user_name" validate:"required,min=3,max=50"`
		NewName    string `json:"user_name,omitempty" validate:"min=3,max=50"`
		IsInactive bool   `json:"is_inactive,omitempty"`
		Recreate   bool   `json:"recreate_token,omitempty"`
	}
)

func NewUserHandler(userI core.UserServiceI) *UserHandler {
	return &UserHandler{
		user: userI,
	}
}

func (h UserHandler) CreateUser(ctx echo.Context) error {

	userReq, err := h.bindUser(ctx)
	if err != nil {
		return ctx.JSON(response.Error(err))
	}

	user := &model.User{
		Name:     userReq.Name,
		IsActive: !userReq.IsInactive,
	}
	if err := h.user.Create(user); err != nil {
		return ctx.JSON(response.Error(err))
	}

	return ctx.JSON(response.OKUser(http.StatusCreated, user))
}

func (h UserHandler) GetUsersList(ctx echo.Context) error {

	users, err := h.user.List()
	if err != nil {
		return ctx.JSON(response.Error(err))
	}

	return ctx.JSON(response.OKUser(http.StatusOK, users))
}

func (h UserHandler) GetUserByName(ctx echo.Context) error {

	userReq, err := h.bindUser(ctx)
	if err != nil {
		return ctx.JSON(response.Error(err))
	}

	user := &model.User{
		Name: userReq.Name,
	}
	if err := h.user.Get(user); err != nil {
		return ctx.JSON(response.Error(err))
	}

	return ctx.JSON(response.OKUser(http.StatusOK, user))
}

func (h UserHandler) UpdateUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, "Not implemented")
	// userReq, err := h.bindUser(ctx)
	// if err != nil {
	// 	return ctx.JSON(response.Error(err))
	// }

	// userName := userReq.Name
	// user := &model.User{
	// 	IsActive: !userReq.IsInactive,
	// }
	// if userReq.NewName != "" && userReq.NewName != userName {
	// 	user.Name = userReq.NewName
	// }
}

func (h UserHandler) bindUser(ctx echo.Context) (UserRequest, error) {

	var user UserRequest

	if err := ctx.Bind(&user); err != nil {
		return UserRequest{}, custom.NewRequestError(err)
	}

	if err := ctx.Validate(&user); err != nil {
		return UserRequest{}, custom.NewRequestError(err)
	}

	user.Name = strings.TrimRight(user.Name, "/")

	return user, nil
}
