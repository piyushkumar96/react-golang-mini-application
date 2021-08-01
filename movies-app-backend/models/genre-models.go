package models

import "time"

// Genre is type for genre
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
}

// MovieGenre is type for movie genre
type MovieGenre struct {
	ID        int       `json:"_"`
	MovieID   int       `json:"_"`
	GenreID   int       `json:"_"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
}
