package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pseudonative/web_page_in_go/pkg/config"
	"github.com/pseudonative/web_page_in_go/pkg/handlers"
	"github.com/pseudonative/web_page_in_go/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	wd, err := os.Getwd()
	if err != nil {
		log.Println("error getting working directory: ", err)
	}

	fmt.Println("Current working directory: ", wd)

	fmt.Printf("starting application on port %s", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
