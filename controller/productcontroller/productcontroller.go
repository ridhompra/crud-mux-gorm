package productcontroller

import (
	"belajar/belajar2/crud-gorm-mux-mysql/helper"
	"belajar/belajar2/crud-gorm-mux-mysql/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

// untuk menampilkan semua data product
func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Products

	if err := models.DB.Find(&products).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, products)
}

// show untuk menampilkan product by id
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// log.Println(vars)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	var product models.Products
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Data tidak ditemukan")
			return
		default:
			ResponseError(w, http.StatusNotFound, err.Error())
			return
		}
	}
	ResponseJson(w, http.StatusOK, product)
	// log.Println(product)
}

// untuk menambah data product
func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Products
	decoder := json.NewDecoder(r.Body)
	//men decode ke variable product
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	//diclose dulu  bodynya
	defer r.Body.Close()
	if err := models.DB.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, product)
	log.Printf("Creating success id :%d", product.Id)
}

// untuk memperbaharui product
func Update(w http.ResponseWriter, r *http.Request) {

}

// untuk menghapus product
func Delete(w http.ResponseWriter, r *http.Request) {

}
