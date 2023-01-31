package routing

import (
	"configer-service/internal/controllers"
	"configer-service/internal/core"
	"configer-service/internal/repositories"
	"configer-service/internal/response"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getIP(req *http.Request) net.IP {
	return net.ParseIP(echo.ExtractIPDirect()(req))
}

func AuthByIP(IPList []net.IPNet) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			ip := getIP(c.Request())

			if ip.IsLoopback() {
				return next(c)
			}

			if len(IPList) > 0 {
				for _, ipnet := range IPList {
					if ipnet.Contains(ip) {
						return next(c)
					}
				}
			}

			return response.ReturnDefault404()
		}
	}
}

func SetRouts(e *echo.Echo, dbConn *gorm.DB) {

	// Working with users
	userRepo := repositories.NewUserRepo(dbConn)
	userService := core.NewUserService(userRepo)
	userHandler := controllers.NewUserHandler(userService)

	user := e.Group("/users")
	var ipList []net.IPNet
	user.Use(AuthByIP(ipList))
	user.POST("/:user_name", userHandler.CreateUser)
	user.GET("/", userHandler.GetUsersList)
	user.GET("/:user_name", userHandler.GetUserByName)

}
