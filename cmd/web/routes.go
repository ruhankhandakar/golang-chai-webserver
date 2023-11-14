package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"ruhan.tech/golang-web/pkg/config"
	"ruhan.tech/golang-web/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(LogRequestInfo)
	mux.Use(NoSurf)
	mux.Use(SetupSession)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	mux.Get("/login", handlers.Repo.LoginHandler)
	mux.Get("/makepost", handlers.Repo.GetMakePostHandler)
	mux.Get("/page", handlers.Repo.PageHandler)
	mux.Get("/article-received", handlers.Repo.ArticleReceivedHandler)

	mux.Post("/makepost", handlers.Repo.PostMakePostHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
