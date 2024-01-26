package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB //It will help the other files to run
)

func Connect() {
	dsn := "root:NO@tcp(localhost:3306)/bookstore_management?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

// func Connect() {
// 	d, err := gorm.Open("mysql", "meghanashree.reddy@localhost?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic(err)
// 	}
// 	db = d // whatever was in d will be transferred to db
// }

func GetDB() *gorm.DB { // talking to db to return the value of db
	return db
}
