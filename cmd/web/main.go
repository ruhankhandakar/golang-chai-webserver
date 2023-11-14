package main

import (
	"encoding/gob"
	"net/http"
	"time"

	"ruhan.tech/golang-web/models"
	"ruhan.tech/golang-web/pkg/config"
	"ruhan.tech/golang-web/pkg/handlers"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager
var app config.AppConfig

func main() {

	gob.Register(models.Article{})

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour // 24 hours
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = false // Dev only, set true in prod
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = sessionManager

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
