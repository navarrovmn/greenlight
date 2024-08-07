package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Models creates a wrapper that will have lots of models
type Models struct {
	Movies      MovieModel
	Permissions PermissionModel
	Tokens      TokenModel
	Users       UserModel
}

// NewModels for ease of us which returns Model struct containing the initialized MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Users:       UserModel{DB: db},
	}
}
