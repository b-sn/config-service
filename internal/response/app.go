package response

import (
	"configer-service/internal/model"

	"net/http"

	"github.com/sirupsen/logrus"
)

type JSONResponseApp struct {
	JSONResponseBase
	App *model.App `json:"app,omitempty"`
}

func OKApp(status int, i any) (int, any) {
	if status >= 400 || http.StatusText(status) == "" {
		logrus.Warnf("wrong HTTP status for success response: %d", status)
	}
	base := JSONResponseBase{
		Status: "OK",
	}
	switch data := i.(type) {
	case *model.App:
		return status, JSONResponseApp{
			JSONResponseBase: base,
			App:              data,
		}
	default:
		logrus.Warnf("unexpected type: %T", data)
		return status, JSONResponseDefault{
			JSONResponseBase: base,
			Data:             data,
		}
	}
}
