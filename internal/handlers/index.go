package handlers

import (
	"html/template"
	"net/http"
)

type IndexHandler struct {
	tmpl *template.Template
}

func NewIndexHandler(t *template.Template) *IndexHandler {
	return &IndexHandler{
		tmpl: t,
	}
}

func (i *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.tmpl.ExecuteTemplate(w, "base", map[string]string{"HelloWorld": "Hello, World!"})
}
