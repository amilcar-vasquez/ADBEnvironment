package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := NewTemplateData()

	data.Title = "Welcome"
	data.HeaderText = "Questions? Comments? Feedback?"

	err := app.render(w, http.StatusOK, "home.tmpl", data)

	if err != nil {
		app.logger.Error("failed to render template", "error", err, "template", "home.tmpl", "url", r.URL.Path, "method", r.Method)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
