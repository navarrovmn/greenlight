package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

// Models creates a wrapper that will have lots of models
type Models struct {
	Movies MovieModel
}

// NewModels for ease of us which returns Model struct containing the initialized MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
