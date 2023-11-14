package render

import (
	"fmt"
	"net/http"
	"text/template"

	"ruhan.tech/golang-web/models"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplate(res http.ResponseWriter, templ string, pgData *models.PageData) {
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
