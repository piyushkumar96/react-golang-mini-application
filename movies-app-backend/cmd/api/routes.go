package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

const apiVersion = "v1"

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//pass httprouter.Params to request context
		ctx := context.WithValue(r.Context(), "params", ps)
		app.logger.Println(ps)
		//call next middleware with new context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	secure := alice.New(app.checkToken)

	// authentication
	router.HandlerFunc(http.MethodPost, "/v1/signin", app.signIn)

	// movies
	router.HandlerFunc(http.MethodGet, "/"+apiVersion+"/movie/:id", app.getMovie)
	router.HandlerFunc(http.MethodGet, "/"+apiVersion+"/movies", app.getAllMovies)

	router.POST("/v1/admin/editmovie", app.wrap(secure.ThenFunc(app.editMovie)))
	// router.HandlerFunc(http.MethodPost, "/"+apiVersion+"/admin/editmovie", app.editMovie)

	// router.DELETE("/v1/admin/movie/:id", app.wrap(secure.ThenFunc(app.deleteMovie)))
	router.DELETE("/v1/admin/movie/:id", app.wrap(secure.ThenFunc(app.deleteMovie)))
	// router.HandlerFunc(http.MethodDelete, "/"+apiVersion+"/admin/movie/:id", app.deleteMovie)

	// genres
	router.HandlerFunc(http.MethodGet, "/"+apiVersion+"/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/"+apiVersion+"/movies/genre/:genre_id", app.getAllMoviesByGenre)
	return app.enableCORS(router)
}
