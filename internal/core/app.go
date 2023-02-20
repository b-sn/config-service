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

const hiddenAppTokenReplace = "*hidden*"

type (
	AppServiceI interface {
		Create(app *model.App) error
		Get(app *model.App) error
	}
	appService struct {
		appRepo   repository.ApplicationI
		appSecret []byte
	}
)

func NewAppService(
	a repository.ApplicationI,
	secret string,
) appService {
	return appService{
		appRepo:   a,
		appSecret: []byte(secret),
	}
}

func (a appService) Create(app *model.App) error {

	findApp := &model.App{
		Name: app.Name,
	}
	if err := a.appRepo.Get(findApp); err != nil {
		LogAppError(findApp, err, "getting application error")
		return errors.New("unexpected get application error")
	}
	if findApp.ID != "" {
		return custom.NewDuplicateError(
			fmt.Errorf(
				"application with name [%s] already exists, try another name",
				findApp.Name,
			),
		)
	}

	if err := a.appRepo.Create(app); err != nil {
		LogAppError(app, err, "creating application error")
		return errors.New("unexpected create application error")
	}

	if token, err := GenerateAppToken(*app, a.appSecret); err != nil {
		return err
	} else {
		app.Token = token
	}

	return nil
}

func (a appService) Get(app *model.App) error {
	if app.Name == "" && app.ID == "" {
		return custom.NewRequestError(
			fmt.Errorf("parameters to get application expected, got: %#v", app),
		)
	}
	if err := a.appRepo.Get(app); err != nil {
		LogAppError(app, err, "getting application error")
		return errors.New("unexpected get application error, see logs")
	}
	return nil
}

func HideAppToken(app *model.App) model.App {
	resApp := *app
	resApp.Token = hiddenAppTokenReplace
	return resApp
}

func LogAppError(app *model.App, err error, msg string) {
	logrus.WithFields(logrus.Fields{
		"app": HideAppToken(app),
	}).Errorf("%s: %w", msg, err)
}

func GenerateAppToken(app model.App, secret []byte) (string, error) {

	payload, err := jwt.NewBuilder().Subject(app.ID).Build()
	if err != nil {
		LogAppError(&app, err, "failed to build token")
		return "", errors.New("unexpected create token error, see logs")
	}

	token, err := jwt.Sign(payload, jwt.WithKey(jwa.HS256, secret))
	if err != nil {
		LogAppError(&app, err, "failed to sign token")
		return "", errors.New("unexpected sign token error, see logs")
	}

	return string(token), nil
}
