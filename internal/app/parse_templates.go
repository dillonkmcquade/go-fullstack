package app

import (
	"html/template"
	"strings"
)

func ParseTemplates() *template.Template {
	// You may add functions for access in your html templates
	funcs := template.FuncMap{
		"toLower": strings.ToLower,
	}
	tmpl := template.Must(template.New("home").Funcs(funcs).ParseGlob("templates/*.html"))
	return tmpl
}
