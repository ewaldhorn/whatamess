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
	// mux.Use(app.enableCORS)

	// authentication routes
	mux.Post("/auth", app.authenticate)
	mux.Post("/refreshToken", app.refreshToken)

	// test handler
	mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		var payload = struct {
			Message string `json:"message"`
		}{
			Message: "hello, world",
		}

		_ = app.writeJSON(w, http.StatusOK, payload)
	})

	// protected routes
	mux.Route("/users", func(mux chi.Router) {
		// use auth middleware

		mux.Get("/", app.allUsers)
		mux.Get("/{userID}", app.retrieveUser)
		mux.Delete("/{userID}", app.deleteUser)
		mux.Put("/", app.createUser)
		mux.Patch("/", app.updateUser)
	})

	return mux
}
