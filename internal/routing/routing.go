package routing

import (
	"configer-service/internal/controllers"
	"configer-service/internal/models"
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

// func AddUser(e echo.Context) error {
// 	return e.JSON(http.StatusNotImplemented, response.ReturnErrorJSON(0, "Adding user is not implemented", nil))
// }

func GetUser(e echo.Context) error {
	return e.JSON(http.StatusNotImplemented, response.ReturnErrorJSON(0, "Getting user is not implemented", nil))
}

func SetRouts(e *echo.Echo, dbConn *gorm.DB) {

	dbConn.AutoMigrate(&models.User{})

	userRepo := repositories.NewUserRepo(dbConn)
	userHandler := controllers.NewUserHandler(userRepo)

	user := e.Group("/users")
	var ipList []net.IPNet
	user.Use(AuthByIP(ipList))
	user.POST("/:user_name", userHandler.AddUser)
	user.GET("/", userHandler.UserList)
	user.GET("/:user_name", GetUser)
}
