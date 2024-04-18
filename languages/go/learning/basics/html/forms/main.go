package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func (c *ContactDetails) toString() string {
	return fmt.Sprintf("Email: '%s', Subject: '%s', Message: '%s'", c.Email, c.Subject, c.Message)
}

func main() {
	tmpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		println("Received", details.toString())

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":9000", nil)
}
