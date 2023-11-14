package handlers

import (
	"net/http"

	"ruhan.tech/golang-web/models"
	"ruhan.tech/golang-web/pkg/config"
	"ruhan.tech/golang-web/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		App: appConfig,
	}
}

func NewHandler(repo *Repository) {
	Repo = repo
}

func (repo *Repository) HomeHandler(res http.ResponseWriter, req *http.Request) {

	repo.App.Session.Put(req.Context(), "userid", "ruhanrk")

	render.RenderTemplate(res, "home.page.html", &models.PageData{})
}
func (repo *Repository) AboutHandler(res http.ResponseWriter, req *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves. We love to talking about ourselves. We are so awesome."

	userId := repo.App.Session.GetString(req.Context(), "userid")
	strMap["userId"] = userId

	render.RenderTemplate(res, "about.page.html", &models.PageData{StrMap: strMap})
}
