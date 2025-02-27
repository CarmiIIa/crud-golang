package productcontroller

import (
	"fmt"
	"golang-crud/entities"
	"golang-crud/models/brandmodel"
	"golang-crud/models/categorymodel"
	"golang-crud/models/productmodel"
	"golang-crud/models/tipemodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	products := productmodel.Getall()

	for i := range products {
		products[i].No = i + 1
	}

	data := map[string]interface{}{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	product := productmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		tipes := tipemodel.GetAll()
		brands := brandmodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"tipes":      tipes,
			"brands":     brands,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		tipeId, err := strconv.Atoi(r.FormValue("tipe_id"))
		if err != nil {
			panic(err)
		}

		brandId, err := strconv.Atoi(r.FormValue("brand_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Tipe.Id = uint(tipeId)
		product.Brand.Id = uint(brandId)
		product.Stock = int64(stock)
		product.Description = r.FormValue("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		// // Handle image upload
		// file, handler, err := r.FormFile("image")
		// if err != nil {
		// 	panic(err)
		// }
		// defer file.Close()

		// // Save the image to a predefined location or upload it to a cloud storage service.
		// // Don't forget to handle naming conflicts and store the image path in the Product struct.
		// imagePath := "/assets/image/" + handler.Filename // Add "/" after "image"
		// fmt.Println("Image Path:", imagePath) // Add this line for debugging
		// f, err := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE, 0666)
		// if err != nil {
		// 	panic(err)
		// }
		// defer f.Close()
		// io.Copy(f, file)

		statusStr := r.FormValue("status")
		fmt.Println("Received statusStr:", statusStr) // Add this line for debugging
		status := entities.Ready
		if statusStr == "TidakReady" {
			status = entities.TidakReady
		}
		product.Status = status

		fmt.Println("Product Status:", product.Status) // Add this line for debugging

		if ok := productmodel.Create(product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		product := productmodel.Detail(id)

		categories := categorymodel.GetAll()
		tipes := tipemodel.GetAll()
		brands := brandmodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"product":    product,
			"tipes":      tipes,
			"brands":     brands,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		tipeId, err := strconv.Atoi(r.FormValue("tipe_id"))
		if err != nil {
			panic(err)
		}

		brandId, err := strconv.Atoi(r.FormValue("brand_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Tipe.Id = uint(tipeId)
		product.Brand.Id = uint(brandId)
		product.Stock = int64(stock)
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()

		if ok := productmodel.Update(id, product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := productmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
