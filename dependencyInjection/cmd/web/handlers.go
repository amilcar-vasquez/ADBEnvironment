package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Let's explore Dependency Injection in Go.\n"))
}
