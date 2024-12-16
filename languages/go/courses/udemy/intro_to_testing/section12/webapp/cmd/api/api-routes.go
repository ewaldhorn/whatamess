package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ----------------------------------------------------------------------------
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// authentication routes
	mux.Post("/auth", app.authenticate)
	mux.Post("/refreshToken", app.refreshToken)

	// protected routes
	mux.Route("/users", func(mux chi.Router) {
		// use auth middleware
		mux.Use(app.authRequired)

		mux.Get("/", app.allUsers)
		mux.Get("/{userID}", app.retrieveUser)
		mux.Delete("/{userID}", app.deleteUser)
		mux.Put("/", app.createUser)
		mux.Patch("/", app.updateUser)
	})

	return mux
}
