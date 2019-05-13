package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Name string
}

func main() {

	templates := template.Must(template.ParseFiles("templates/index.html"))
	p := Page{Name: "Gopher"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
