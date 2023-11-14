package main

import (
	"net/http"

	"ruhan.tech/golang-web/pkg/config"
	"ruhan.tech/golang-web/pkg/handlers"
)

func main() {
	var app config.AppConfig

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	serve := &http.Server{
		Addr:    ":8081",
		Handler: routes(&app),
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
