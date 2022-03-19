package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>home page</h1>")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>contact page</h1>")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>login page</h1>")
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>register page</h1>")
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Page not found</h1>")
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
