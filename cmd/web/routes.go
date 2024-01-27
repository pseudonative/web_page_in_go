package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/pseudonative/web_page_in_go/pkg/config"
	"github.com/pseudonative/web_page_in_go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
