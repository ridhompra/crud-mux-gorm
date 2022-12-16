package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	db, err := gorm.Open(mysql.Open(DbMaster))
	if err != nil {
		log.Println("Connection DB Failed")
	}
	db.AutoMigrate(&Products{})
	DB = db
}
