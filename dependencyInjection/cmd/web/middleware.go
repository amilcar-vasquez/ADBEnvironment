package main

import (
	"net/http"
)

func (a *application) loggingMiddleware(next http.Handler) http.Handler {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// gather the log values before creating the log message to avoid overcrowding
		// the Printf() function
		var (
			ip     = r.RemoteAddr
			proto  = r.Proto
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		a.logger.Info("receive request", "remote_addr", ip, "proto", proto, "method", method, "uri", uri)
		// pre-processing
		a.logger.Info("Pre-Processing")
		next.ServeHTTP(w, r)
		// post-processing
		a.logger.Info("Post-Processing")
	})
	return fn

}
