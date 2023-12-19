package app

import (
	"html/template"
	"strings"
)

// Reads all templates from templates/*.html, loads a FuncMap, and returns the template
func Templates() *template.Template {
	// FuncMap is a mapping of functions that can be used inside templates
	funcs := template.FuncMap{
		"toLower": strings.ToLower,
	}
	// ParseGlob is parsing all templates with the .html extension in the templates dir, this can be changed if needed
	tmpl := template.Must(template.New("home").Funcs(funcs).ParseGlob("templates/*.html"))
	return tmpl
}
