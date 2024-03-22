package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

const pathToTemplates = "./templates/"
const pathToPages = pathToTemplates + "pages"
const baseLayoutTemplate = "base.layout.gohtml"

// ------------------------------------------------------------------------------------------------
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

// ------------------------------------------------------------------------------------------------
type TemplateData struct {
	IP   string
	Data map[string]any
}

// ------------------------------------------------------------------------------------------------
func (app *application) render(w http.ResponseWriter, r *http.Request, pageTemplate string, data *TemplateData) error {
	parsedTemplate, err := template.ParseFiles(path.Join(pathToPages, pageTemplate), path.Join(pathToPages, baseLayoutTemplate))
	if err != nil {
		http.Error(w, "Error - Bad Request", http.StatusBadRequest)
		println(err.Error())
		return err
	}

	data.IP = app.ipFromContext(r.Context())

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

// ------------------------------------------------------------------------------------------------
func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request (login)", http.StatusBadRequest)
		return
	}

	form := NewForm(r.PostForm)
	form.Required("email", "password")

	if !form.Valid() {
		fmt.Fprintf(w, "failed validation")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println(email, password)

	fmt.Fprintf(w, "We have [%s]:[%s]", email, password)
}
