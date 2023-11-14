package handlers

import (
	"log"
	"net/http"

	"ruhan.tech/golang-web/models"
	"ruhan.tech/golang-web/pkg/config"
	"ruhan.tech/golang-web/pkg/forms"
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

	render.RenderTemplate(res, req, "home.page.html", &models.PageData{})
}

func (repo *Repository) AboutHandler(res http.ResponseWriter, req *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves. We love to talking about ourselves. We are so awesome."

	userId := repo.App.Session.GetString(req.Context(), "userid")
	strMap["userId"] = userId

	render.RenderTemplate(res, req, "about.page.html", &models.PageData{StrMap: strMap})
}

func (repo *Repository) LoginHandler(res http.ResponseWriter, req *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(res, req, "login.page.html", &models.PageData{StrMap: strMap})
}

func (repo *Repository) GetMakePostHandler(res http.ResponseWriter, req *http.Request) {

	var emptyArticle models.Article
	data := make(map[string]interface{})
	data["article"] = emptyArticle

	render.RenderTemplate(res, req, "make-post.page.html", &models.PageData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (repo *Repository) PostMakePostHandler(res http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	blogTitle := req.Form.Get("blog_title")
	blogArticle := req.Form.Get("blog_article")

	article := models.Article{
		BlogTitle:   blogTitle,
		BlogArticle: blogArticle,
	}

	form := forms.New(req.PostForm)

	form.HasRequired("blog_title", "blog_article")

	form.MinLength("blog_title", 5, req)
	form.MinLength("blog_article", 5, req)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["article"] = article

		render.RenderTemplate(res, req, "make-post.page.html", &models.PageData{
			Form: form,
			Data: data,
		})
		return
	}

	repo.App.Session.Put(req.Context(), "article", article)
	http.Redirect(res, req, "/article-received", http.StatusSeeOther)

}

func (repo *Repository) ArticleReceivedHandler(res http.ResponseWriter, req *http.Request) {
	article, ok := repo.App.Session.Get(req.Context(), "article").(models.Article)

	if !ok {
		log.Println("Can't get data from session")

		repo.App.Session.Put(req.Context(), "error", "Can't get data from session")

		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)

		return
	}

	data := make(map[string]interface{})
	data["article"] = article

	render.RenderTemplate(res, req, "article-received.page.html",
		&models.PageData{
			Data: data,
		})
}

func (repo *Repository) PageHandler(res http.ResponseWriter, req *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(res, req, "page.page.html", &models.PageData{StrMap: strMap})
}
