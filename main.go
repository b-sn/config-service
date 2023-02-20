package main

import (
	"configer-service/internal/custom"
	"configer-service/internal/db"
	"configer-service/internal/routing"
	"configer-service/internal/structs"
	"configer-service/pkg/utils"
	"flag"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	cfgFileName := flag.String("cfg", "config-dev", "the name of the config file")
	flag.Parse()

	viper.SetConfigName(*cfgFileName)
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		// TODO: Separate config not found error
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	cfg := structs.CfgData{}
	viper.Unmarshal(&cfg)

	fmt.Println(utils.PrettyPrint(cfg))

	e := echo.New()

	e.Validator = custom.NewCustomValidator()

	// Set number of requests per second
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	// e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())
	// e.Use(middleware.Secure())
	// e.Use(middleware.Logger())

	if cfg.Env == "test" {
		os.Remove(cfg.DB.File)
	}

	dbConn := db.GetSQLiteConnection(cfg.DB.File, &gorm.Config{})

	routing.SetRouts(e, dbConn, cfg.Security)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
