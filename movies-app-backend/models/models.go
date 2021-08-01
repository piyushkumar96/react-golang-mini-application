package models

import "database/sql"

// Models is wrapper for database
type Models struct {
	DB DBModels
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModels{
			DB: db,
		},
	}
}
