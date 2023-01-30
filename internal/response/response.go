package response

import (
	"configer-service/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
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
	User models.User `json:"user,omitempty"`
}

type JSONResponseUsers struct {
	JSONResponseBase
	Users []models.User `json:"users"`
}

func ReturnOkJSON(i interface{}) interface{} {
	base := JSONResponseBase{
		Status: "OK",
	}
	switch data := i.(type) {
	case models.User:
		return JSONResponseUser{
			JSONResponseBase: base,
			User:             data,
		}
	case []models.User:
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

func ReturnErrorJSON(errCode int, errMsg string) interface{} {
	response := JSONResponseBase{
		Status: "Error",
	}
	if errCode != 0 {
		response.ErrorCode = errCode
	}
	if errMsg != "" {
		response.Message = errMsg
	}
	return response
}

func ReturnDefault404() error {
	return echo.NewHTTPError(http.StatusNotFound)
}

func ReturnDefault400() error {
	return echo.NewHTTPError(http.StatusBadRequest)
}
