package models

import "ruhan.tech/golang-web/pkg/forms"

type PageData struct {
	StrMap          map[string]string
	IntMap          map[string]int
	FltMap          map[string]float32
	DataMap         map[string]interface{}
	CSRFToken       string
	Warning         string
	Error           string
	IsAuthenticated int
	Form            *forms.Form
	Data            map[string]interface{}
}
