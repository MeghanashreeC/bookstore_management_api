package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB //It will help the other files to run
)

func Connect() { // Connecting to db
	dsn := "root:123456@tcp(0.0.0.0:3306)/bookstore"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB { // talking to db to return the value of db
	return db
}
