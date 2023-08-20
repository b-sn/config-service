package response

import (
	"configer-service/internal/model"
	"net/http"

	"github.com/sirupsen/logrus"
)

type JSONResponseConfig struct {
	JSONResponseBase
	Config *model.Config `json:"config,omitempty"`
}

func OKConfig(status int, i any) (int, any) {
	if status >= 400 || http.StatusText(status) == "" {
		logrus.Warnf("wrong HTTP status for success response: %d", status)
	}
	base := JSONResponseBase{
		Status: "OK",
	}
	switch data := i.(type) {
	case *model.Config:
		return status, JSONResponseConfig{
			JSONResponseBase: base,
			Config:           data,
		}
	default:
		logrus.Warnf("unexpected type: %T", data)
		return status, JSONResponseDefault{
			JSONResponseBase: base,
			Data:             data,
		}
	}
}
