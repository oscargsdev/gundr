package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/statc", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /projects/{id}", app.projectView)

	return mux
}
