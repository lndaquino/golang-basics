package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", handleSite)
	http.HandleFunc("/home", handleHome)
	http.HandleFunc("/user", handleUser)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func handleSite(w http.ResponseWriter, r *http.Request) {
	u := user{"João", "joao@gmail.com"}
	templates.ExecuteTemplate(w, "home.html", u)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página user!"))
}
