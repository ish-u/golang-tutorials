package config

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ORM Object
var (
	db *gorm.DB
)

// Connect Function
func Connect() {
	// d, err := gorm.Open("mysql", "(root:admin)/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	db = d
}

// Returning the DB
func GetDB() *gorm.DB {
	return db
}
