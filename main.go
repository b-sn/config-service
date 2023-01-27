package main

import (
	"configer-service/internal/routing"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	// Set number of requests per second
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())

	dbConn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	routing.SetRouts(e, dbConn)

	e.Logger.Fatal(e.Start(":1323"))
}
