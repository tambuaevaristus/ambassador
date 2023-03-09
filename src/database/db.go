package database

import "gorm.io/gorm"
import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
  )
var DB *gorm.DB

func Connect(){
	_, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}
}

func AutoMigrate(){
	DB.AutoMigrate(models.User{})
}