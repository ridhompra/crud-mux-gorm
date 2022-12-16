package router

import (
	"belajar/belajar2/crud-gorm-mux-mysql/controller/productcontroller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	port := ":8080"
	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/products/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/products", productcontroller.Create).Methods("POST")
	r.HandleFunc("/products/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/products", productcontroller.Delete).Methods("DELETE")

	log.Printf("Server Running port %s \n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
