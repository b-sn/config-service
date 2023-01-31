package main

import (
	"configer-service/internal/db"
	"configer-service/internal/routing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	// Set number of requests per second
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	// e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())
	// e.Use(middleware.Secure())
	// e.Use(middleware.Logger())

	dbConn := db.GetSQLiteConnection("gorm.db", &gorm.Config{})

	routing.SetRouts(e, dbConn)

	e.Logger.Fatal(e.Start(":1323"))
}
