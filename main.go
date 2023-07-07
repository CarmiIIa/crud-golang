package main

import (
	"golang-crud/config"
	"golang-crud/controllers/categorycontroller"
	"golang-crud/controllers/homecontroller"
	"golang-crud/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/create", categorycontroller.Create)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/create", productcontroller.Create)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
