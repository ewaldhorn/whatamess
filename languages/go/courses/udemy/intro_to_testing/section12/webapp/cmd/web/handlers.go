package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
	"webapp/pkg/data"
)

// need to keep as vars, since tests need to adjust the relative path location(s)
var pathToTemplates = "./templates/"
var pathToPages = pathToTemplates + "pages"
var pathToProfilePics = "./static/img"

const baseLayoutTemplate = "base.layout.gohtml"
const MAX_UPLOAD_SIZE int64 = 1024 * 1024 * 5

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

	if app.Session.Exists(r.Context(), "user") {
		templateData.User = app.Session.Get(r.Context(), "user").(data.User)
	}

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
		http.Redirect(w, r, HOME_URL, http.StatusSeeOther)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.DB.GetUserByEmail(email)
	if err != nil {
		// redirect to login page with an error message
		app.Session.Put(r.Context(), "error", "Login failed.")
		http.Redirect(w, r, HOME_URL, http.StatusSeeOther)
		return
	}

	// authenticate password
	if !app.authenticate(r, user, password) {
		app.Session.Put(r.Context(), "error", "Login failed.")
		http.Redirect(w, r, HOME_URL, http.StatusSeeOther)
		return
	}

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

// ----------------------------------------------------------------------------
func (app *application) authenticate(r *http.Request, user *data.User, password string) bool {
	if valid, err := user.PasswordMatches(password); err != nil || !valid {
		return false
	}

	app.Session.Put(r.Context(), "user", user)
	return true
}

// ----------------------------------------------------------------------------
func (app *application) UploadProfilePic(w http.ResponseWriter, r *http.Request) {
	// get file from request
	files, err := app.UploadFiles(r, pathToProfilePics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get user from session
	user := app.Session.Get(r.Context(), "user").(data.User)

	// create a data.UserImage variable
	var newUserImage = data.UserImage{
		UserID:   user.ID,
		FileName: files[0].OriginalFilename,
	}

	// insert image into database
	_, err = app.DB.InsertUserImage(newUserImage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// refresh user in session since there's a new profile image
	updatedUser, err := app.DB.GetUser(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	app.Session.Put(r.Context(), "user", updatedUser)

	// redirect to profile page
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
}

// ----------------------------------------------------------------------------
type UploadedFile struct {
	OriginalFilename string
	FileSize         int64
}

// ----------------------------------------------------------------------------
func (app *application) UploadFiles(r *http.Request, uploadDir string) ([]*UploadedFile, error) {
	var uploadedFiles []*UploadedFile

	err := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
	if err != nil {
		return nil, fmt.Errorf("the uploaded file is too big and must be under %d bytes", MAX_UPLOAD_SIZE)
	}

	for _, fHeaders := range r.MultipartForm.File {
		for _, hdr := range fHeaders {
			uploadedFiles, err = func(uploadedFiles []*UploadedFile) ([]*UploadedFile, error) {
				var uploadedFile UploadedFile

				infile, err := hdr.Open()
				if err != nil {
					return nil, err
				}
				defer infile.Close()

				uploadedFile.OriginalFilename = hdr.Filename

				var outfile *os.File
				defer outfile.Close()

				if outfile, err = os.Create(filepath.Join(uploadDir, uploadedFile.OriginalFilename)); err != nil {
					return nil, err
				} else {
					fileSize, err := io.Copy(outfile, infile)
					if err != nil {
						return nil, err
					}
					uploadedFile.FileSize = fileSize
				}
				uploadedFiles = append(uploadedFiles, &uploadedFile)
				return uploadedFiles, nil
			}(uploadedFiles)

			if err != nil {
				return uploadedFiles, err
			}
		}
	}

	return uploadedFiles, nil
}
