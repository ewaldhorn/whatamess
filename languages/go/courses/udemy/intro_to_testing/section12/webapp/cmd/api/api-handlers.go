package main

import (
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// ----------------------------------------------------------------------------
type Credentals struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

// ----------------------------------------------------------------------------
func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	var credentials Credentals
	// read json payload
	err := app.readJSON(w, r, &credentials)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// look up user by email
	user, err := app.DB.GetUserByEmail(credentials.Username)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		app.errorJSON(w, errors.New("authorized"), http.StatusUnauthorized)
		return
	}

	// generate tokens
	tokens, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// return tokens
	err = app.writeJSON(w, http.StatusOK, tokens)
	if err != nil {
		log.Printf("error returning tokens %v", err)
	}
}

// ----------------------------------------------------------------------------
func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {

}

// ----------------------------------------------------------------------------
func (app *application) allUsers(w http.ResponseWriter, r *http.Request) {

}

// ----------------------------------------------------------------------------
func (app *application) retrieveUser(w http.ResponseWriter, r *http.Request) {

}

// ----------------------------------------------------------------------------
func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {

}

// ----------------------------------------------------------------------------
func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {

}

// ----------------------------------------------------------------------------
func (app *application) createUser(w http.ResponseWriter, r *http.Request) {

}
