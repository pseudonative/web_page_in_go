package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		http.Error(w, "Internal Server Error one", http.StatusInternalServerError)
		return
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

// main is the main application function
func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println("error getting working directory: ", err)
	}
	fmt.Println("Current working directory: ", wd)

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Printf("starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
