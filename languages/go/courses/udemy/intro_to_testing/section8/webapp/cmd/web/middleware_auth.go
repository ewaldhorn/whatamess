package main

import "net/http"

// ----------------------------------------------------------------------------
func (app *application) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.Session.Exists(r.Context(), "user") {
			app.Session.Put(r.Context(), "error", "Not logged in")
			http.Redirect(w, r, HOME_URL, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
