package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/justinas/nosurf"
)

func LogRequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		// Log the request info
		now := time.Now()
		fmt.Printf("%d/%d/%d - %d:%d\n", now.Month(), now.Day(), now.Year(), now.Hour(), now.Minute())
		fmt.Println("Path: ", req.URL.Path)

		next.ServeHTTP(res, req)
	})
}

func SetupSession(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}

func NoSurf(next http.Handler) http.Handler {
	noSurfHandler := nosurf.New(next)
	noSurfHandler.SetBaseCookie(http.Cookie{
		Name:     "mycsrfcookie",
		Path:     "/",
		Domain:   "",
		Secure:   false, // Dev only, set true in prod
		HttpOnly: true,
		MaxAge:   3600,
		SameSite: http.SameSiteLaxMode,
	})

	return noSurfHandler
}
