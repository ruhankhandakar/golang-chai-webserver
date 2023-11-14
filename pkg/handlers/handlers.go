package handlers

import (
	"net/http"

	"ruhan.tech/golang-web/models"
	"ruhan.tech/golang-web/pkg/render"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "home.page.html", &models.PageData{})
}
func AboutHandler(res http.ResponseWriter, req *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves. We love to talking about ourselves. We are so awesome."

	render.RenderTemplate(res, "about.page.html", &models.PageData{StrMap: strMap})
}
