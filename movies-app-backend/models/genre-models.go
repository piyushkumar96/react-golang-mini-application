package models

import "time"

// Genre is type for genre
type Genre struct {
	ID        int       `json:"id,omitempty"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// MovieGenre is type for movie genre
type MovieGenre struct {
	ID        int       `json:"id,omitempty"`
	MovieID   int       `json:"movie_id,omitempty"`
	GenreID   int       `json:"genre_id,omitempty"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
