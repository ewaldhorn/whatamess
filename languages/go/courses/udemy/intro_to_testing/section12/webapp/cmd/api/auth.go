package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// ----------------------------------------------------------------------------
const jwtTokenExpiry = time.Minute * 15
const refreshTokenExpirty = time.Hour * 24

// ----------------------------------------------------------------------------
type TokenPair struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// ----------------------------------------------------------------------------
type Claims struct {
	UserName string `json:"name"`
	jwt.RegisteredClaims
}

// ----------------------------------------------------------------------------
func (app *application) getAndVerifyTokenFromHeader(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {
	// add header
	w.Header().Add("Vary", "Authorization")

	// read auth header
	authHeader := r.Header.Get("Authorization")

	// sanity check
	if authHeader == "" {
		return "", nil, errors.New("no auth header")
	}

	// split header into components
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", nil, errors.New("invalid auth header")
	}

	// check for Bearer
	if headerParts[0] != "Bearer" {
		return "", nil, errors.New("unauthorized: Bearer missing")
	}

	token := headerParts[1]
	claims := &Claims{}

	// parse token
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// validate signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(app.JWTSecret), nil
	})

	// check for errors or expired tokens
	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired by") {
			return "", nil, errors.New("expired token")
		}

		return "", nil, err
	}

	// make sure we issued this token
	if claims.Issuer != app.Domain {
		return "", nil, errors.New("incorrect issuer")
	}

	// finally, return valid token
	return token, claims, nil
}
