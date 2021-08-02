package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DBModels struct {
	DB *sql.DB
}

// Get returns one movie and error if any
func (m *DBModels) GetMovie(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := fmt.Sprintf(`select 
				id, title, description, year, release_date, runtime, rating, cbfc_rating, created_at, updated_at 
			  from movies 
			  where id = $1`)
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.Rating,
		&movie.CBFCRating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	genreQuery := fmt.Sprintf(`select
				 mg.id, mg.movie_id, mg.genre_id, g.genre_name 
			 from 
			 	movies_genres mg 
			 left join
			 	genres g 
			 on (mg.genre_id = g.id)
			 where mg.movie_id = $1
			`)

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
func (m *DBModels) GetAllMovies(genre ...int) ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(genre) > 0 {
		where = fmt.Sprintf("where id in (select movie_id from movies_genres where genre_id = %d)", genre[0])
	}
	query := fmt.Sprintf(`select 
				id, title, description, year, release_date, runtime, rating, cbfc_rating, created_at, updated_at 
			  from movies
			  %s	
			  order by title
			  `, where)

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
			&movie.Rating,
			&movie.CBFCRating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genreQuery := fmt.Sprintf(`select
				 mg.id, mg.movie_id, mg.genre_id, g.genre_name 
			 from 
			 	movies_genres mg 
			 left join
			 	genres g 
			 on (mg.genre_id = g.id)
			 where mg.movie_id = $1
			`)

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

func (m *DBModels) InsertMovie(movie Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`insert into movies 
			 (title, description, year, release_date, runtime, rating, cbfc_rating, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			`)

	_, err := m.DB.ExecContext(ctx, stmt, movie.Title,
		movie.Description, movie.Year, movie.ReleaseDate, movie.Runtime, movie.Rating, movie.CBFCRating, movie.CreatedAt, movie.UpdatedAt)

	if err != nil {
		return err
	}
	return err
}

func (m *DBModels) UpdateMovie(movie Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`update movies set
			 title = $1, description = $2, year = $3, release_date = $4, runtime = $5, rating = $6, cbfc_rating = $7, updated_at = $8 
			where id = $9
			`)

	_, err := m.DB.ExecContext(ctx, stmt, movie.Title,
		movie.Description, movie.Year, movie.ReleaseDate, movie.Runtime, movie.Rating, movie.CBFCRating, movie.UpdatedAt, movie.ID)

	if err != nil {
		return err
	}
	return nil
}

func (m *DBModels) DeleteMovie(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`delete from movies where id = $1`)

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
