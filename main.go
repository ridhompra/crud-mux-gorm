package main

import (
	"belajar/belajar2/crud-gorm-mux-mysql/models"
	"belajar/belajar2/crud-gorm-mux-mysql/router"
)

func main() {
	models.ConnectionDB()
	router.Router()
}
