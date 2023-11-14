package main

import (
	"fmt"
	"net/http"
	"time"
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
