package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err:= gorm.Open(mysql.Open("root:@tcp(localhost:3306)/golang_relational"))

	if err != nil {
		panic(err)
	}

	DB = db
}