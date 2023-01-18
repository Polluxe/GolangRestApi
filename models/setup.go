package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("golangAPI:golang@tcp(182.16.245.106:3306)/golangapi"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&item{})

	DB = database
}
