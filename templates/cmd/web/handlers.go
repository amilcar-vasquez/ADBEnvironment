package main

import (
	"github.com/cohune-cabbage/di/internal/data"
	"github.com/cohune-cabbage/di/internal/validator"
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

func (app *application) createFeedback(w http.ResponseWriter, r *http.Request) {
	//parse the form data
	err := r.ParseForm()
	if err != nil {
		app.logger.Error("failed to parse form", "error", err)
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}
	//return a map. the names are from the template
	name := r.Form.Get("name")
	email := r.PostForm.Get("email")
	subject := r.PostForm.Get("subject")
	message := r.PostForm.Get("message")

	//create a feedback instance
	feedback := &data.Feedback{
		FullName: name,
		Email:    email,
		Subject:  subject,
		Message:  message,
	}

	//validate the feedback
	v := validator.NewValidator()
	data.ValidateFeedback(v, feedback)
	if !v.ValidData() {
		data := NewTemplateData()
		data.Title = "Welcome"
		data.HeaderText = "Questions? Comments? Feedback?"
		data.FormErrors = v.Errors
		data.FormData = map[string]string{
			"name":    name,
			"email":   email,
			"subject": subject,
			"message": message,
		}
		err := app.render(w, http.StatusUnprocessableEntity, "home.tmpl", data)
		if err != nil {
			app.logger.Error("failed to render template", "error", err, "template", "home.tmpl", "url", r.URL.Path, "method", r.Method)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}

	err = app.feedback.Insert(feedback)
	if err != nil {
		app.logger.Error("failed to insert feedback", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
