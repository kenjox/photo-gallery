package main

import (
	"fmt"
	"net/http"
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

/**
	Creating custom router which implements ServeHTTP(ResponseWriter, *Request)
	since ListenAndServe
**/

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homePage(w, r)
	case "/contact":
		contactPage(w, r)
	case "/login":
		loginPage(w, r)
	case "/register":
		registerPage(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Server running on port 9000")
	http.ListenAndServe(":9000", router)
}
