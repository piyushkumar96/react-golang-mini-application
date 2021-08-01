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

	// genres
	router.HandlerFunc(http.MethodGet, "/"+apiVersion+"/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/"+apiVersion+"/movies/genre/:genre_id", app.getAllMoviesByGenre)
	return app.enableCORS(router)
}
