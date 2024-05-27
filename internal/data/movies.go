package data

import (
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`                // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`                 // Timestamp for when it's added to the DB
	Title     string    `json:"title"`             // Movie title
	Year      int32     `json:"year,omitempty"`    // Movie release year
	Runtime   int32     `json:"runtime,omitempty"` // Runtime in minutes
	Genres    []string  `json:"genres,omitempty"`  // Slices of genres for the movie
	Version   int32     `json:"version"`           // Version number starts at 1 and will be incremented whenever movie is updated
}
