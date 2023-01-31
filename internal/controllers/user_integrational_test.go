package controllers

import (
	"configer-service/internal/core"
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

var handler *UserHandler
var testToken string

const testUserName = "test_user"

func TestMain(m *testing.M) {
	const test_db_file = "test.db"
	os.Remove(test_db_file)

	dbConn := db.GetSQLiteConnection(test_db_file, &gorm.Config{})
	userRepo := repositories.NewUserRepo(dbConn)
	userService := core.NewUserService(userRepo)
	handler = NewUserHandler(userService)

	// userRepo := repositories.NewUserRepo(dbConn)
	// handler = NewUserHandler(userRepo)

	code := m.Run()

	sql, _ := dbConn.DB()
	sql.Close()
	os.Remove(test_db_file)

	os.Exit(code)
}

func getContextAndRecorder(method string) (echo.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(
		httptest.NewRequest(method, "/", nil),
		rec,
	)
	return ctx, rec
}

func TestCreateUser(t *testing.T) {
	ctx, rec := getContextAndRecorder(http.MethodPost)
	ctx.SetParamNames("user_name")
	ctx.SetParamValues(testUserName)

	// Assertions
	if assert.NoError(t, handler.CreateUser(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var respData response.JSONResponseUser
		json.Unmarshal(rec.Body.Bytes(), &respData)
		assert.Equal(t, "OK", respData.Status)

		user := respData.User
		assert.Equal(t, testUserName, user.Name)

		assert.NotEmpty(t, user.Token)
		testToken = user.Token
		assert.True(t, user.IsActive)
		assert.Empty(t, user.ID, "Expected not to return user ID")
	}

	// if assert.NoError(t, handler.CreateUser(ctx)) {
	// 	assert.Equal(t, http.StatusConflict, rec.Code)

	// 	var respData response.JSONResponseDefault
	// 	json.Unmarshal(rec.Body.Bytes(), &respData)
	// 	assert.Equal(t, "Error", respData.Status)
	// }
}

func TestGetUserByName(t *testing.T) {
	ctx, rec := getContextAndRecorder(http.MethodGet)
	ctx.SetParamNames("user_name")
	ctx.SetParamValues(testUserName)

	if assert.NoError(t, handler.GetUserByName(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var respData response.JSONResponseUser
		json.Unmarshal(rec.Body.Bytes(), &respData)
		assert.Equal(t, "OK", respData.Status)

		user := respData.User
		assert.Equal(t, testToken, user.Token)
		assert.True(t, user.IsActive)
		assert.Empty(t, user.ID, "Expected not to return user ID")
	}
}

func TestUserList(t *testing.T) {
	ctx, rec := getContextAndRecorder(http.MethodGet)
	if assert.NoError(t, handler.GetUsersList(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var respData response.JSONResponseUsers
		json.Unmarshal(rec.Body.Bytes(), &respData)

		users := respData.Users
		assert.Equal(t, 1, len(users))

		user := users[0]
		assert.Equal(t, "*hidden*", user.Token)
		assert.Equal(t, testUserName, user.Name)
		assert.Empty(t, user.ID, "Expected not to return user ID")
	}
}
