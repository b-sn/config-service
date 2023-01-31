package db

import (
	"configer-service/internal/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSQLiteConnection(fileName string, cfg *gorm.Config) *gorm.DB {
	dbConn, err := gorm.Open(sqlite.Open(fileName), cfg)
	if err != nil {
		log.Fatal(err)
	}

	dbConn.AutoMigrate(&models.User{})

	return dbConn
}
