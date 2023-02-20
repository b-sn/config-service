package routing

import (
	"configer-service/internal/controller"
	"configer-service/internal/core"
	"configer-service/internal/repository"
	"configer-service/internal/structs"
	"configer-service/pkg/echoutils"
	"fmt"
	"net"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetRouts(e *echo.Echo, dbConn *gorm.DB, cfg structs.Secure) {

	// Working with users
	userRepo := repository.NewUserRepo(dbConn)
	userService := core.NewUserService(userRepo, cfg.UserTokenSecret)
	userHandler := controller.NewUserHandler(userService)

	user := e.Group("/users")

	var ipList []net.IPNet
	if len(cfg.AllowedIP) > 0 {
		for _, ip := range cfg.AllowedIP {
			_, ipnet, err := net.ParseCIDR(ip)
			if err != nil {
				panic(fmt.Sprintf("Error in IP [%s]: %v", ip, err))
			}
			ipList = append(ipList, *ipnet)
		}
	}
	user.Use(echoutils.AuthByIP(ipList))
	user.POST("/:user_name", userHandler.CreateUser)
	user.GET("/", userHandler.GetUsersList)
	user.GET("/:user_name", userHandler.GetUserByName)
	user.PATCH("/:user_name", userHandler.UpdateUser)

	// Working with application
	appRepo := repository.NewAppRepo(dbConn)
	appService := core.NewAppService(appRepo, userRepo, cfg.AppTokenSecret)
	appHandler := controller.NewAppHandler(appService)

	app := e.Group("/applications")
	app.Use(echoutils.AuthorizeUser([]byte(cfg.UserTokenSecret)))
	app.POST("/:app_name", appHandler.CreateApp)
	app.GET("/:app_name", appHandler.GetAppByName)

	// Working with config
	cfgRepo := repository.NewConfigRepo(dbConn)
	cfgService := core.NewConfigService(cfgRepo)
	cfgHandler := controller.NewConfigHandler(cfgService)

	config := e.Group("/config")
	config.Use(echoutils.AuthorizeUser([]byte(cfg.UserTokenSecret)))
	config.POST("/:app_name/:env_name", cfgHandler.CreateConfig)
}
