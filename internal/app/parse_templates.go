package app

import "html/template"

func ParseTemplates() *template.Template {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	return tmpl
}
