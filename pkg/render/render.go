package render

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/justinas/nosurf"
	"ruhan.tech/golang-web/models"
)

var templateCache = make(map[string]*template.Template)

func AddCSRFData(pgData *models.PageData, req *http.Request) *models.PageData {
	pgData.CSRFToken = nosurf.Token(req)
	return pgData
}

func RenderTemplate(res http.ResponseWriter, req *http.Request, templ string, pgData *models.PageData) {
	var tmpl *template.Template
	var err error

	_, inMap := templateCache[templ]

	if !inMap {
		err = makeTemplateCache(templ)
		if err != nil {
			fmt.Println("Error creating template cache", err)
		}
	} else {
		fmt.Println("Template already in cache")
	}
	tmpl = templateCache[templ]

	pgData = AddCSRFData(pgData, req)

	err = tmpl.Execute(res, pgData)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

}

func makeTemplateCache(templ string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", templ),
		"./templates/base.layout.html",
	}

	parsedTemplate, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}
	templateCache[templ] = parsedTemplate
	return nil
}
