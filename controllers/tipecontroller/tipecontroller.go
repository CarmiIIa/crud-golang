package tipecontroller

import (
	"golang-crud/entities"
	"golang-crud/models/tipemodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	tipes := tipemodel.GetAll()
	data := map[string]any {
		"tipes": tipes,
	}

	temp, err := template.ParseFiles("views/tipe/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/tipe/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var tipe entities.Tipe

		tipe.Name = r.FormValue("name")
		tipe.CreatedAt = time.Now()
		tipe.UpdatedAt = time.Now()

		if ok := tipemodel.Create(tipe); !ok {
			temp, _ := template.ParseFiles("views/tipe/create.html")
			temp.Execute(w, nil)
		}
		
		http.Redirect(w, r, "/tipe", http.StatusSeeOther)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/tipe/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		tipe := tipemodel.Detail(id)
		data := map[string]any{
			"tipe": tipe,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var tipe entities.Tipe
	
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
	
		tipe.Name = r.FormValue("name")
		tipe.UpdatedAt = time.Now()
	
		ok, err := tipemodel.Update(id, tipe)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if !ok {
			temp, err := template.ParseFiles("views/tipe/edit.html")
			if err != nil {
				http.Error(w, "Failed to render template", http.StatusInternalServerError)
				return
			}
			temp.Execute(w, nil)
			return
		}
		
		http.Redirect(w, r, "/tipe", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := tipemodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/tipe", http.StatusSeeOther)
}