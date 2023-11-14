package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"ruhan.tech/golang-web/pkg/config"
	"ruhan.tech/golang-web/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)

	return mux
}
