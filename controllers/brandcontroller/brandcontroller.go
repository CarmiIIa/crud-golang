package brandcontroller

import (
	"golang-crud/entities"
	"golang-crud/models/brandmodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	brands := brandmodel.GetAll()
	data := map[string]any {
		"brands": brands,
	}

	temp, err := template.ParseFiles("views/brand/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/brand/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var brand entities.Brand

		brand.Name = r.FormValue("name")
		brand.CreatedAt = time.Now()
		brand.UpdatedAt = time.Now()

		if ok := brandmodel.Create(brand); !ok {
			temp, _ := template.ParseFiles("views/brand/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/brand", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/brand/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		brand := brandmodel.Detail(id)
		data := map[string]any{
			"brand": brand,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var brand entities.Brand

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		brand.Name = r.FormValue("name")
		brand.UpdatedAt = time.Now()

		ok, err := brandmodel.Update(id, brand)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !ok {
			temp, err := template.ParseFiles("views/brand/edit.html")
			if err != nil {
				http.Error(w, "Failed to render template", http.StatusInternalServerError)
				return
			}
			temp.Execute(w, nil)
			return
		}

		http.Redirect(w, r, "/brand", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := brandmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/brand", http.StatusSeeOther)
}