package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModels struct {
	DB *sql.DB
}

// Get returns one movie and error if any
func (m *DBModels) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select 
				id, title, description, year, release_date, runtime, cbfc_rating, created_at, updated_at 
			  from movies 
			  where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.CBFCRating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	genreQuery := `select
				 mg.id, mg.movie_id, mg.genre_id, g.genre_name 
			 from 
			 	movies_genres mg 
			 left join
			 	genres g 
			 on (mg.genre_id = g.id)
			 where mg.movie_id = $1
			`

	rows, err := m.DB.QueryContext(ctx, genreQuery, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	genres := make(map[int]string)

	for rows.Next() {
		var mg MovieGenre
		err = rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}

		genres[mg.ID] = mg.Genre.GenreName
	}
	movie.MovieGenre = genres

	return &movie, nil
}

// GetAll returns all movies and error if any
func (m *DBModels) GetAll() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select 
				id, title, description, year, release_date, runtime, cbfc_rating, created_at, updated_at 
			  from movies
			  `

	rows, err := m.DB.QueryContext(ctx, query)
	// defer rows.Close()
	if err != nil {
		return nil, err
	}

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.CBFCRating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genreQuery := `select
				 mg.id, mg.movie_id, mg.genre_id, g.genre_name 
			 from 
			 	movies_genres mg 
			 left join
			 	genres g 
			 on (mg.genre_id = g.id)
			 where mg.movie_id = $1
			`

		genreRows, err := m.DB.QueryContext(ctx, genreQuery, movie.ID)
		if err != nil {
			return nil, err
		}

		genres := make(map[int]string)

		for genreRows.Next() {
			var mg MovieGenre
			err = genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}

			genres[mg.ID] = mg.Genre.GenreName
		}
		genreRows.Close()

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}
