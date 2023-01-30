package controllers

import (
	"configer-service/internal/models"
	"configer-service/internal/response"
	"log"
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

func (h *UserHandler) CreateUser(e echo.Context) error {
	var user *models.User
	if err := h.bindUser(e, user); err != nil {
		return err
	}
	if err := h.userRepo.Create(user); err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			response.ReturnErrorJSON(createUserError, "User creation error"))
	}
	return e.JSON(http.StatusCreated, response.ReturnOkJSON(user)) // TODO: Check, probably here should be *user
}

func (h *UserHandler) UserList(e echo.Context) error {
	users, err := h.userRepo.List()
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			response.ReturnErrorJSON(createUserError, "Getting user list error"))
	}
	return e.JSON(http.StatusOK, response.ReturnOkJSON(users))
}

func (h *UserHandler) GetUserByName(e echo.Context) error {
	var user models.User
	if err := h.bindUser(e, &user); err != nil {
		// log.Fatalln(err)
		return err
	}
	// log.Fatalln(user)
	users, err := h.userRepo.GetUserByName(user.Name)
	if err != nil {
		return e.JSON(
			http.StatusInternalServerError,
			response.ReturnErrorJSON(createUserError, "Getting user list error"))
	}
	log.Fatalln(users)

	return e.JSON(http.StatusOK, response.ReturnOkJSON(users[0]))
}

func (h *UserHandler) bindUser(e echo.Context, user *models.User) error {
	// log.Println("0000001")
	// var u models.User
	if err := e.Bind(user); err != nil {
		// log.Println("0000002")
		return e.JSON(
			http.StatusBadRequest,
			response.ReturnErrorJSON(paramsError, "Param user_name error"))
	}
	// log.Println(user)
	// log.Println("0000003")
	return nil
}
