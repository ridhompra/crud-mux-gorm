package main

import (
	"crud-gorm-mux-mysql/models"
	"crud-gorm-mux-mysql/router"
)

func main() {
	models.ConnectionDB()
	router.Router()
}
