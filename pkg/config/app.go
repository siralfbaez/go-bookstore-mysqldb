package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

//DdsRMyFav34

func Connect() {
	d, err := gorm.Open("mysql", "siralf:DdsRMyFav34@tcp(127.0.0.1:3306)/justagodb")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
