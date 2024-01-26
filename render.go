package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + html)
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
