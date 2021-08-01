package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const apiVersion = "v1"

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	// movies
	router.HandlerFunc(http.MethodGet, "/"+apiVersion+"/movie/:id", app.getMovie)
	router.HandlerFunc(http.MethodGet, "/"+apiVersion+"/movies", app.getAllMovies)
	return app.enableCORS(router)
}
