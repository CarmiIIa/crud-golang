package main

import (
	"golang-crud/config"
	"golang-crud/controllers/brandcontroller"
	"golang-crud/controllers/categorycontroller"
	"golang-crud/controllers/homecontroller"
	"golang-crud/controllers/productcontroller"
	"golang-crud/controllers/tipecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

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

	// tipe
	http.HandleFunc("/tipe", tipecontroller.Index)
	http.HandleFunc("/tipe/create", tipecontroller.Create)
	http.HandleFunc("/tipe/edit", tipecontroller.Edit)
	http.HandleFunc("/tipe/delete", tipecontroller.Delete)

	// brand
	http.HandleFunc("/brand", brandcontroller.Index)
	http.HandleFunc("/brand/create", brandcontroller.Create)
	http.HandleFunc("/brand/edit", brandcontroller.Edit)
	http.HandleFunc("/brand/delete", brandcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
