package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pseudonative/web_page_in_go/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println("error getting working directory: ", err)
	}
	fmt.Println("Current working directory: ", wd)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
