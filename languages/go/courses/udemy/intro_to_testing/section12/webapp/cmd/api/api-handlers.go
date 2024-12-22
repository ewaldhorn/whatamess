package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// ----------------------------------------------------------------------------
type Credentals struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

// ----------------------------------------------------------------------------
func (app *application) defaultRoute(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, "Yellow there!")
	if err != nil {
		log.Printf("error returning tokens %v", err)
	}
}

// ----------------------------------------------------------------------------
func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	var credentials Credentals
	// read json payload
	err := app.readJSON(w, r, &credentials)
	if err != nil {
		app.errorJSON(w, errors.New("bad request"), http.StatusBadRequest)
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
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	refreshToken := r.Form.Get("refresh_token")
	claims := &Claims{}

	_, err = jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.JWTSecret), nil
	})
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) > 30*time.Second {
		app.errorJSON(w, errors.New("refresh token does not need to be renewed yet"), http.StatusTooEarly)
		return
	}

	// time to try to process the request
	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.DB.GetUser(userID)
	if err != nil {
		app.errorJSON(w, errors.New("unknown user"), http.StatusBadRequest)
		return
	}

	tokens, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "__Host-refresh_token",
		Path:     "/",
		Value:    tokens.RefreshToken,
		Expires:  time.Now().Add(refreshTokenExpiry),
		MaxAge:   int(refreshTokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   true,
	})

	err = app.writeJSON(w, http.StatusOK, tokens)
	if err != nil {
		log.Println(err.Error())
	}
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
