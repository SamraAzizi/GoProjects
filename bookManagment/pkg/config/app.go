package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db = gorm.DB
)

func connect() {
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?")
	if err != nil {
		panic(err)

	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
