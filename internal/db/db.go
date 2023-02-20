package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSQLiteConnection(fileName string, cfg *gorm.Config) *gorm.DB {
	dbConn, err := gorm.Open(sqlite.Open(fileName), cfg)
	if err != nil {
		log.Fatal(err)
	}

	return dbConn
}
