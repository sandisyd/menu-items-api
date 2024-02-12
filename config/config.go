package config

import (
	"menu/menuItems/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/menu_db?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(models.MenuItems{}, models.Submenu{})

	DB = db
}
