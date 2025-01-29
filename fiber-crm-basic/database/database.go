package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("lead.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
}
