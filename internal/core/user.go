package core

import (
	"configer-service/internal/custom"
	"configer-service/internal/models"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
)

const hiddenTokenReplace = "*hidden*"

type UserServiceI interface {
	Create(*models.User) error
	List() ([]*models.User, error)
	Get(*models.User) error
}

type userService struct {
	repo models.UserRepository
}

func NewUserService(r models.UserRepository) userService {
	return userService{
		repo: r,
	}
}

func (u userService) Create(user *models.User) error {
	if user.Name == "" {
		return custom.RequestError{
			Text: fmt.Sprintf("name expected to create user, got: %#v", user),
		}
	}
	if err := u.repo.Get(user); err != nil {
		log.Errorf("getting user error: %w, user: %#v", err, user)
		return errors.New("unexpected get user error, see logs")
	}
	if user.ID != 0 {
		return custom.DuplicateError{
			Text: fmt.Sprintf("User with name [%s] already exists, try another name", user.Name),
		}
	}
	if err := u.repo.Create(user); err != nil {
		log.Errorf("creating user error: %w, user: %#v", err, user)
		return errors.New("unexpected create user error, see logs")
	}
	return nil
}

func (u userService) List() ([]*models.User, error) {
	users, err := u.repo.Find(&models.User{})
	if err != nil {
		log.Errorf("getting users list error: %w", err)
		return nil, errors.New("unexpected user list error, see logs")
	}
	for _, user := range users {
		user.Token = hiddenTokenReplace
	}
	return users, nil
}

func (u userService) Get(user *models.User) error {
	if user.Name == "" && user.ID == 0 && user.Token == "" {
		return custom.RequestError{
			Text: fmt.Sprintf("parameters to get user expected, got: %#v", user),
		}
	}
	if err := u.repo.Get(user); err != nil {
		log.Errorf("getting user error: %w, user: %#v", err, user)
		return errors.New("unexpected get user error, see logs")
	}
	return nil
}
