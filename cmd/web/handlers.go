package main

import (
	"crypto/rand"
	"html/template"
	"net/http"
	"net/url"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, "parsing template", err)
		return
	}

	templateData := map[string]any{
		"error": app.sessionManager.PopString(r.Context(), "formErrors"),
		"flash": app.sessionManager.PopString(r.Context(), "flash"),
		"urls":  app.memStore,
	}

	if err := t.ExecuteTemplate(w, "base", templateData); err != nil {
		app.serverError(w, r, "executing template", err)
		return
	}
}

func (app *application) shortUrlCreate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.serverError(w, r, "reading post form", err)
		return
	}

	formError := ""

	longUrl := r.PostForm.Get("long_url")
	if longUrl == "" {
		formError = "must not be empty"
	} else if _, err := url.ParseRequestURI(longUrl); err != nil {
		formError = "must be a valid url"
	}

	if formError != "" {
		app.sessionManager.Put(r.Context(), "formErrors", formError)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	shortUrl := rand.Text()[:8]
	for _, ok := app.memStore[shortUrl]; ok; {
		shortUrl = rand.Text()[:8]
	}

	app.memStore[shortUrl] = longUrl
	app.logger.Info("successfully created short url", "short url", shortUrl)

	app.sessionManager.Put(r.Context(), "flash", "Short URL successfully created!")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) redirectUrl(w http.ResponseWriter, r *http.Request) {
	url, ok := app.memStore[r.PathValue("url")]
	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}
