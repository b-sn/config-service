package response

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type JSONResponse struct {
	Status    string      `json:"status"`
	ErrorCode int         `json:"error_code,omitempty"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func ReturnOkJSON(data interface{}) (int, JSONResponse) {
	return http.StatusOK, JSONResponse{
		Status: "OK",
		Data:   data,
	}
}

func ReturnErrorJSON(errCode int, errMsg string, err error) JSONResponse {
	if err != nil {
		fmt.Printf("Error code: %d, message %s, source %v\n", errCode, errMsg, err)
	}
	response := JSONResponse{
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
