package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, html string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+html, "./templates/base.layout.html")
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

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	// check to see if we already have template in cache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// template in cache
		log.Println("using cached template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}
	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache
	tc[t] = tmpl
	return nil
}
