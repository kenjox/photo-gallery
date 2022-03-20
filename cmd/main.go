package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

// getPageName get the name of the file without extension e.g home.gohtml ==> home
func getPageName(file string) string {
	return strings.Split(file, ".")[0]
}

// loadTemplate loads the template using the file name
func loadTemplate(w http.ResponseWriter, templateName string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	file := filepath.Join("templates", templateName)
	tpl, err := template.ParseFiles(file)

	if err != nil {
		log.Printf("Error parsing %v", err)
		pageName := fmt.Sprintf("Error loading %s page", getPageName(templateName))
		http.Error(w, pageName, http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)

	if err != nil {
		log.Printf("Error executing %v", err)
		pageName := fmt.Sprintf("Error loading %s page", getPageName(templateName))
		http.Error(w, pageName, http.StatusInternalServerError)
		return
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	loadTemplate(w, "home.gohtml")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	loadTemplate(w, "contact.gohtml")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	loadTemplate(w, "login.gohtml")
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	loadTemplate(w, "register.gohtml")
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	loadTemplate(w, "404.gohtml")
}

func main() {
	r := chi.NewRouter()

	r.Get("/", homePage)
	r.Get("/contact", contactPage)
	r.Get("/register", registerPage)
	r.Get("/login", loginPage)
	r.NotFound(notFoundPage)

	fmt.Println("Server running on port 9000")
	http.ListenAndServe(":9000", r)
}
