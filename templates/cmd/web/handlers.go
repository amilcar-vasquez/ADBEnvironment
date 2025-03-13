package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	page := "home.tmpl"
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("The template %s does not exist", page)
		app.logger.Error("Template does not exist", "page", page, "error", err, "url", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":      "Let's explore Dependency Injection in Go.",
		"HeaderText": "Welcome to the Dependency Injection in Go.",
	}

	err := ts.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
