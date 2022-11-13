package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil)) //Servidor já está rodando
}

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{
		"João",
		"joao.pedro@gmail.com",
	}

	templates = template.Must(template.ParseGlob("*.html"))
	templates.ExecuteTemplate(w, "home.html", u)
}
