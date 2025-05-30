package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

// need to keep as vars, since tests need to adjust the relative path location(s)
var pathToTemplates = "./templates/"
var pathToPages = pathToTemplates + "pages"

const baseLayoutTemplate = "base.layout.gohtml"

// ----------------------------------------------------------------------------
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var templateData = make(map[string]any)

	if app.Session.Exists(r.Context(), "test") {
		msg := app.Session.GetString(r.Context(), "test")
		templateData["test"] = msg
	} else {
		app.Session.Put(r.Context(), "test", "Hit this page at "+time.Now().UTC().String())
	}

	_ = app.render(w, r, "home.page.gohtml", &TemplateData{Data: templateData})
}

// ----------------------------------------------------------------------------
type TemplateData struct {
	IP   string
	Data map[string]any
}

// ----------------------------------------------------------------------------
func (app *application) render(
	w http.ResponseWriter,
	r *http.Request,
	pageTemplate string,
	data *TemplateData) error {
	parsedTemplate, err := template.ParseFiles(path.Join(pathToPages, pageTemplate),
		path.Join(pathToPages, baseLayoutTemplate))
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

// ----------------------------------------------------------------------------
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
