package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
	"webapp/pkg/data"
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
	IP    string
	Data  map[string]any
	Error string
	Flash string
	User  data.User
}

// ----------------------------------------------------------------------------
func (app *application) render(
	w http.ResponseWriter,
	r *http.Request,
	pageTemplate string,
	templateData *TemplateData) error {
	parsedTemplate, err := template.ParseFiles(path.Join(pathToPages, pageTemplate),
		path.Join(pathToPages, baseLayoutTemplate))
	if err != nil {
		http.Error(w, "Error - Bad Request", http.StatusBadRequest)
		println(err.Error())
		return err
	}

	templateData.IP = app.ipFromContext(r.Context())
	templateData.Error = app.Session.PopString(r.Context(), "error")
	templateData.Flash = app.Session.PopString(r.Context(), "flash")

	err = parsedTemplate.Execute(w, templateData)
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
		// redirect to login page with an error message
		app.Session.Put(r.Context(), "error", "Login failed.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.DB.GetUserByEmail(email)
	if err != nil {
		// redirect to login page with an error message
		app.Session.Put(r.Context(), "error", "Login failed.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// authenticate password
	fmt.Println(user, password)

	// prevent fixation attack
	_ = app.Session.RenewToken(r.Context())

	// store success in Session
	app.Session.Put(r.Context(), "flash", "Login succeeded.")
	// redirect to app
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)

}

// ----------------------------------------------------------------------------
func (app *application) Profile(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "profile.page.gohtml", &TemplateData{})
}
