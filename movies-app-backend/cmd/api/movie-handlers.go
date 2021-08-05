package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"net/url"
	"react-golang-mini-application/movies-app-backend/models"
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
	if movie.Poster == "" {
		movie = getPoster(movie)
	}

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
	params := r.Context().Value("params").(httprouter.Params)
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

func getPoster(movie models.Movie) models.Movie {
	type TheMovieDB struct {
		Page    int `json:"page"`
		Results []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			GenreIds         []int   `json:"genre_ids"`
			ID               int     `json:"id"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			Popularity       float64 `json:"popularity"`
			PosterPath       string  `json:"poster_path"`
			ReleaseDate      string  `json:"release_date"`
			Title            string  `json:"title"`
			Video            bool    `json:"video"`
			VoteAverage      float64 `json:"vote_average"`
			VoteCount        int     `json:"vote_count"`
		} `json:"results"`
		TotalPages   int `json:"total_pages"`
		TotalResults int `json:"total_results"`
	}

	client := &http.Client{}
	apiKey := "77d435fa382fc6698e146b209f7b5c77"

	theURL := "https://api.themoviedb.org/3/search/movie?api_key="
	req, err := http.NewRequest("GET", theURL+apiKey+"&query="+url.QueryEscape(movie.Title), nil)
	if err != nil {
		log.Println(err)
		return movie
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return movie
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return movie
	}

	var responseObject TheMovieDB

	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		log.Println(err)
		return movie
	}

	if len(responseObject.Results) > 0 {
		movie.Poster = responseObject.Results[0].PosterPath
	}

	return movie
}
