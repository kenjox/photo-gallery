package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	HTMLTempl *template.Template
}

func Parse(filePath string) (Template, error) {
	tpl, err := template.ParseFiles(filePath)

	if err != nil {
		return Template{}, fmt.Errorf("error parsing %v", err)
	}

	return Template{
		HTMLTempl: tpl,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.HTMLTempl.Execute(w, data)
	if err != nil {
		log.Printf("error executing %v", err)
		http.Error(w, "Error excuting template", http.StatusInternalServerError)
		return
	}
}
