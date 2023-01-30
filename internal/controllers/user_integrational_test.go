package controllers

import (
	"configer-service/internal/db"
	"configer-service/internal/repositories"
	"configer-service/internal/response"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const test_db_file = "test.db"

func TestCreateUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users/test_user", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	dbConn := db.GetSQLiteConnection(test_db_file, &gorm.Config{})
	defer func() {
		sql, _ := dbConn.DB()
		sql.Close()
		os.Remove(test_db_file)
	}()

	userRepo := repositories.NewUserRepo(dbConn)
	handler := NewUserHandler(userRepo)

	// Assertions
	if assert.NoError(t, handler.CreateUser(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var respData response.JSONResponseUser
		json.Unmarshal(rec.Body.Bytes(), &respData)
		// assert.Contains(t, respData, "user")
		user := respData.User
		// data := models.User{respData.Data}

		assert.Equal(t, "test_user", user.Name)

		// assert.Equal(t, userJSON, rec.Body.String())
		// assert.
	}
}
