package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"movies-app-backend/models"
	"net/http"
	"strconv"
	"time"
)

type jsonResp struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (app *application) getMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	movie, err := app.models.DB.GetMovie(id)

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.GetAllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) editMovie(w http.ResponseWriter, r *http.Request) {
	var payload models.MoviePayLoad
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var movie models.Movie

	if payload.ID != "0" {
		id, err := strconv.Atoi(payload.ID)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		m, err := app.models.DB.GetMovie(id)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		movie = *m
	}

	movie.ID, _ = strconv.Atoi(payload.ID)
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.ReleaseDate, _ = time.Parse("2006-01-02", payload.ReleaseDate)
	movie.Year = movie.ReleaseDate.Year()
	movie.Runtime, _ = strconv.Atoi(payload.Runtime)
	movie.Rating, _ = strconv.Atoi(payload.Rating)
	movie.CBFCRating = payload.CBFCRating
	movie.UpdatedAt = time.Now()

	if movie.ID == 0 {
		movie.CreatedAt = time.Now()
		err = app.models.DB.InsertMovie(movie)
	} else {
		err = app.models.DB.UpdateMovie(movie)
	}
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{
		Ok: true,
	}
	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.DB.DeleteMovie(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{
		Ok: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
