package main

import (
	"log"
	"os"
	"testing"
	"time"
	"webapp/pkg/repository/databaserepo"

	"github.com/golang-jwt/jwt/v4"
)

// ----------------------------------------------------------------------------
var app application

// ----------------------------------------------------------------------------
func TestMain(m *testing.M) {
	app.DB = &databaserepo.TestingDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "ee334ab190919293949a9b9cc1c2c3c4c5c6c7c8c910"

	os.Exit(m.Run())
}

// ----------------------------------------------------------------------------
func createValidToken() string {
	// generate a token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["sub"] = "1"
	claims["admin"] = true
	claims["aud"] = "example.com"
	claims["iss"] = "example.com"
	// leave this to 3 days, for easy manual testing
	expires := time.Now().UTC().Add(time.Hour * 72)
	claims["exp"] = expires.Unix()

	signedAccessToken, err := token.SignedString([]byte(app.JWTSecret))
	if err != nil {
		log.Fatal(err)
	}
	return signedAccessToken
}

// ----------------------------------------------------------------------------
func createExpiredToken() string {
	// generate a token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["sub"] = "1"
	claims["admin"] = true
	claims["aud"] = "example.com"
	claims["iss"] = "example.com"
	// make sure the token is already expired
	expires := time.Now().UTC().Add(time.Hour * 72 * -1)
	claims["exp"] = expires.Unix()

	signedAccessToken, err := token.SignedString([]byte(app.JWTSecret))
	if err != nil {
		log.Fatal(err)
	}
	return signedAccessToken
}
