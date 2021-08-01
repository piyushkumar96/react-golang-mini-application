package models

import "time"

// Movie is type for movies
type Movie struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Year        int            `json:"year"`
	ReleaseDate time.Time      `json:"release_date"`
	Runtime     int            `json:"runtime"`
	Rating      int            `json:"rating"`
	CBFCRating  string         `json:"cbfc_rating"`
	CreatedAt   time.Time      `json:"_"`
	UpdatedAt   time.Time      `json:"_"`
	MovieGenre  map[int]string `json:"genres"`
}
