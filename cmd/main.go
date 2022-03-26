package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kenjox/photo-gallery/views"
)

// loadTemplate loads the template using the file name
func loadTemplate(w http.ResponseWriter, filePath string) {
	tpl, err := views.Parse(filePath)

	if err != nil {
		log.Printf("error occurred loading %v", err)
		http.Error(w, "unable to load the template", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	loadTemplate(w, "templates/home.gohtml")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	loadTemplate(w, "templates/contact.gohtml")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	loadTemplate(w, "templates/login.gohtml")
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	loadTemplate(w, "templates/register.gohtml")
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	loadTemplate(w, "templates/404.gohtml")
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
