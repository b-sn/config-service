package response

import (
	"configer-service/internal/custom"
	"configer-service/internal/models"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type JSONResponseBase struct {
	Status    string `json:"status"`
	ErrorCode int    `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
}

type JSONResponseDefault struct {
	JSONResponseBase
	Data interface{} `json:"data,omitempty"`
}

type JSONResponseUser struct {
	JSONResponseBase
	User *models.User `json:"user,omitempty"`
}

type JSONResponseUsers struct {
	JSONResponseBase
	Users []*models.User `json:"users"`
}

func ReturnOkJSON(i interface{}) interface{} {
	base := JSONResponseBase{
		Status: "OK",
	}
	switch data := i.(type) {
	case *models.User:
		return JSONResponseUser{
			JSONResponseBase: base,
			User:             data,
		}
	case []*models.User:
		return JSONResponseUsers{
			JSONResponseBase: base,
			Users:            data,
		}
	default:
		return JSONResponseDefault{
			JSONResponseBase: base,
			Data:             data,
		}
	}
}

func getStatus(err error) int {
	switch err.(type) {

	case custom.DuplicateError:
		return http.StatusConflict

	case custom.RequestError:
		return http.StatusBadRequest

	case custom.NotFoundError:
		return http.StatusNotFound

	default:
		return http.StatusInternalServerError

	}
}

func getErrorMessage(err error) string {
	if err == nil || err.Error() == "" {
		return http.StatusText(getStatus(err))
	}
	return err.Error()
}

func Error(e echo.Context, err error) error {
	return e.JSON(
		getStatus(err),
		JSONResponseBase{
			Status:  "Error",
			Message: getErrorMessage(err),
		},
	)
}

func OK(e echo.Context, status int, i interface{}) error {
	if status >= 400 || http.StatusText(status) == "" {
		log.Errorf("wrong HTTP status for success response: %d", status)
	}
	base := JSONResponseBase{
		Status: "OK",
	}
	switch data := i.(type) {
	case *models.User:
		return e.JSON(status, JSONResponseUser{
			JSONResponseBase: base,
			User:             data,
		})
	case []*models.User:
		return e.JSON(status, JSONResponseUsers{
			JSONResponseBase: base,
			Users:            data,
		})
	default:
		return e.JSON(status, JSONResponseDefault{
			JSONResponseBase: base,
			Data:             data,
		})
	}
}

func ReturnDefault404() error {
	return echo.NewHTTPError(http.StatusNotFound)
}

func ReturnDefault400() error {
	return echo.NewHTTPError(http.StatusBadRequest)
}
