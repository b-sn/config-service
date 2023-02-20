package response

import (
	"configer-service/internal/model"

	"net/http"

	"github.com/sirupsen/logrus"
)

type JSONResponseUser struct {
	JSONResponseBase
	User *model.User `json:"user,omitempty"`
}

type JSONResponseUsers struct {
	JSONResponseBase
	Users []*model.User `json:"users"`
}

func OKUser(status int, i any) (int, any) {
	if status >= 400 || http.StatusText(status) == "" {
		logrus.Warnf("wrong HTTP status for success response: %d", status)
	}
	base := JSONResponseBase{
		Status: "OK",
	}
	switch data := i.(type) {
	case *model.User:
		return status, JSONResponseUser{
			JSONResponseBase: base,
			User:             data,
		}
	case []*model.User:
		return status, JSONResponseUsers{
			JSONResponseBase: base,
			Users:            data,
		}
	default:
		logrus.Warnf("unexpected type: %T", data)
		return status, JSONResponseDefault{
			JSONResponseBase: base,
			Data:             data,
		}
	}
}
