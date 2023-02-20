package response

import (
	"configer-service/internal/custom"
	"net/http"

	"github.com/sirupsen/logrus"
)

type JSONResponseBase struct {
	Status    string `json:"status"`
	ErrorCode int    `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
}

type JSONResponseDefault struct {
	JSONResponseBase
	Data any `json:"data,omitempty"`
}

func GetStatus(err error) int {
	switch err.(type) {

	case custom.DuplicateError:
		return http.StatusConflict

	case custom.RequestError:
		return http.StatusBadRequest

	case custom.UnauthorizedError:
		return http.StatusUnauthorized

	case custom.NotFoundError:
		return http.StatusNotFound

	default:
		logrus.Warnf("error with unknown type returned: %w", err)
		return http.StatusInternalServerError
	}
}

func GetErrorMessage(err error) string {
	if err == nil || err.Error() == "" {
		return http.StatusText(GetStatus(err))
	}
	return err.Error()
}

func Error(err error) (int, JSONResponseBase) {
	return GeneralError(GetStatus(err), GetErrorMessage(err))
}

func GeneralError(status int, msg string) (int, JSONResponseBase) {
	return status, JSONResponseBase{
		Status:  "Error",
		Message: msg,
	}
}

func UnauthorizedError(msg string) (int, JSONResponseBase) {
	return GeneralError(http.StatusUnauthorized, msg)
}
