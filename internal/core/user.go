package core

import (
	"configer-service/internal/custom"
	"configer-service/internal/model"
	"configer-service/internal/repository"
	"errors"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/sirupsen/logrus"
)

const hiddenUserTokenReplace = "*hidden*"

type UserServiceI interface {
	Create(*model.User) error
	List() ([]*model.User, error)
	Get(*model.User) error
}

type userService struct {
	repo       repository.UserI
	userSecret []byte
}

func NewUserService(r repository.UserI, secret string) userService {
	return userService{
		repo:       r,
		userSecret: []byte(secret),
	}
}

func (u userService) Create(user *model.User) error {

	if userExists, err := u.repo.Exists(user.Name); err != nil {
		LogUserError(user, err, "getting user error")
		return errors.New("unexpected get user error, see logs")
	} else if userExists {
		return custom.NewDuplicateError(
			fmt.Errorf("user with name [%s] already exists, try another name", user.Name),
		)
	}

	if err := u.repo.Create(user); err != nil {
		LogUserError(user, err, "creating user error")
		return errors.New("unexpected create user error, see logs")
	}

	if token, err := GenerateUserToken(*user, u.userSecret); err != nil {
		return err
	} else {
		user.Token = token
	}

	return nil
}

func (u userService) List() ([]*model.User, error) {
	users, err := u.repo.Find(&model.User{})
	if err != nil {
		logrus.Errorf("getting users list error: %w", err)
		return nil, errors.New("unexpected user list error, see logs")
	}

	for i := 0; i < len(users); i++ {
		newUser := HideUserToken(users[i])
		users[i] = &newUser
	}
	return users, nil
}

func (u userService) Get(user *model.User) error {
	if user.Name == "" && user.ID == "" {
		return custom.NewRequestError(
			fmt.Errorf("parameters to get user expected, got: %#v", user),
		)
	}
	if err := u.repo.Get(user); err != nil {
		LogUserError(user, err, "getting user error")
		return errors.New("unexpected get user error, see logs")
	}

	// TODO: Make separate method for creating new token
	if token, err := GenerateUserToken(*user, u.userSecret); err != nil {
		return err
	} else {
		user.Token = token
	}

	return nil
}

func HideUserToken(user *model.User) model.User {
	resUser := *user
	resUser.Token = hiddenUserTokenReplace
	return resUser
}

func LogUserError(user *model.User, err error, msg string) {
	logrus.WithFields(logrus.Fields{
		"user": HideUserToken(user),
	}).Errorf("%s: %w", msg, err)
}

func GenerateUserToken(user model.User, secret []byte) (string, error) {

	payload, err := jwt.NewBuilder().Subject(user.ID).Build()
	if err != nil {
		LogUserError(&user, err, "failed to build token")
		return "", errors.New("unexpected create token error, see logs")
	}

	token, err := jwt.Sign(payload, jwt.WithKey(jwa.HS256, secret))
	if err != nil {
		LogUserError(&user, err, "failed to sign token")
		return "", errors.New("unexpected sign token error, see logs")
	}

	return string(token), nil
}
