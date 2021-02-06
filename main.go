package main

import (
	"os"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// port := os.Getenv("PORT")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/stories", storiesHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/main.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "main", nil)
}

func storiesHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/stories.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "stories", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/about.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "about", nil)
}