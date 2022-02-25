// Code generated by entc, DO NOT EDIT.

package movie

import (
	"time"
)

const (
	// Label holds the string label denoting the movie type in the database.
	Label = "movie"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldRated holds the string denoting the rated field in the database.
	FieldRated = "rated"
	// FieldRealeaseDate holds the string denoting the realease_date field in the database.
	FieldRealeaseDate = "realease_date"
	// FieldGenre holds the string denoting the genre field in the database.
	FieldGenre = "genre"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldPoster holds the string denoting the poster field in the database.
	FieldPoster = "poster"
	// Table holds the table name of the movie in the database.
	Table = "movies"
)

// Columns holds all SQL columns for movie fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldYear,
	FieldRated,
	FieldRealeaseDate,
	FieldGenre,
	FieldLanguage,
	FieldPoster,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultYear holds the default value on creation for the "year" field.
	DefaultYear string
	// DefaultRealeaseDate holds the default value on creation for the "realease_date" field.
	DefaultRealeaseDate func() time.Time
)
