package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /r/{url}", app.redirectUrl)
	mux.HandleFunc("POST /url/create", app.shortUrlCreate)

	return app.logRequest(app.sessionManager.LoadAndSave((mux)))
}
