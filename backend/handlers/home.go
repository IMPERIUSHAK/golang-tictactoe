package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var home = template.Must(template.ParseFiles(filepath.Join("..", "templates", "login_page.html")))

func LoginPage(w http.ResponseWriter, r *http.Request) {

	home, err := template.ParseFiles(filepath.Join("..", "templates", "login_page.html"))
	if err != nil {
		http.Error(w, "Error while loading login page", http.StatusInternalServerError)
	}

	if err := home.Execute(w, ""); err != nil {
		http.Error(w, "Error while loading login page", http.StatusInternalServerError)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	home, err := template.ParseFiles(filepath.Join("..", "templates", "home.html"))
	if err != nil {
		http.Error(w, "Error while loading Home page", http.StatusInternalServerError)
	}

	if err := home.Execute(w, ""); err != nil {
		http.Error(w, "Error while loading home page", http.StatusInternalServerError)
	}
}
