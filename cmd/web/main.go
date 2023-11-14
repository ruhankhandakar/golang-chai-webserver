package main

import (
	"log"
	"net/http"

	"ruhan.tech/golang-web/pkg/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	err := http.ListenAndServe(":8081", nil)
	log.Fatal(err)
}
