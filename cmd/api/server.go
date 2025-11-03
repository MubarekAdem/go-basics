package main

import (
	"fmt"
	"net/http"
)
func serve(app *application) error {
	srver := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
		idleTimeout: time.minute,
		readTimeout: 10 * time.Second,
		writeTimeout: 30 * time.Second,
	}
	log.Printf("starting %s server on port %d", app.config.env, app.config.port)
	return srver.ListenAndServe()
}