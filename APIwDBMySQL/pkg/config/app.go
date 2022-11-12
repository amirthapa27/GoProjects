package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// establishing a connection to the database
func Connect() {
	// dsn := "amir:mypass/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", "root:mypass@/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
