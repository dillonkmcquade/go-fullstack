package handlers

import (
	"html/template"
	"log"
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
	var err error
	err = i.tmpl.ExecuteTemplate(w, "base", map[string]string{"HelloWorld": "Hello, World!"})
	if err != nil {
		log.Printf("executing template: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
