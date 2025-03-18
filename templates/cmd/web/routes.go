package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("POST /feedback", app.createFeedback)
	mux.HandleFunc("GET /feedback/success", app.feedbackSuccess)

	return app.loggingMiddleware(mux)
}

func (app *application) feedbackSuccess(w http.ResponseWriter, r *http.Request) {
	data := NewTemplateData()
	data.Title = "Feedback Submitted"
	data.HeaderText = "Thank you for your feedback!"

	err := app.render(w, http.StatusOK, "feedback_success.tmpl", data)

	if err != nil {
		app.logger.Error("failed to render feedback success page", "error", err, "template", "feedback_success.tmpl", "url", r.URL.Path, "method", r.Method)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
