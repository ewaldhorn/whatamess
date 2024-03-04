package main

import (
	"net/http"
	"text/template"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "pages/home.page.gohtml", &TemplateData{})
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, pageTemplate string, data *TemplateData) error {
	parsedTemplate, err := template.ParseFiles("./templates/" + pageTemplate)
	if err != nil {
		http.Error(w, "Error - Bad Request", http.StatusBadRequest)
		return err
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
