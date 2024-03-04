package main

import (
	"html/template"
	"net/http"
	"path"
)

var pathToTemplates = "./templates/"

// ------------------------------------------------------------------------------------------------
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "pages/home.page.gohtml", &TemplateData{})
}

// ------------------------------------------------------------------------------------------------
type TemplateData struct {
	IP   string
	Data map[string]any
}

// ------------------------------------------------------------------------------------------------
func (app *application) render(w http.ResponseWriter, r *http.Request, pageTemplate string, data *TemplateData) error {
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, pageTemplate))
	if err != nil {
		http.Error(w, "Error - Bad Request", http.StatusBadRequest)
		return err
	}

	data.IP = app.ipFromContext(r.Context())

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
