package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

var pathToTemplates = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	err := app.render(w, r, "home.page.gohtml", &TemplateData{})
	if err != nil {
		log.Print("Error rendering home page", err)
	}
}

func (app *application) About(w http.ResponseWriter, r *http.Request) {
	err := app.render(w, r, "about.page.gohtml", &TemplateData{})
	if err != nil {
		log.Print("Error rendering about page", err)
	}
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// read and parse the Go template
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	data.IP = app.ipFromContext(r.Context())

	// now execute the template with the passed data
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return err
	}

	return nil
}
