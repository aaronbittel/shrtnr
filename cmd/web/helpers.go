package main

import (
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, context string, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	if app.debug {
		app.logger.Error(context, "method", method, "uri", uri, "err", err, "stack", debug.Stack())
	} else {
		app.logger.Error(context, "method", method, "uri", uri, "err", err)
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
