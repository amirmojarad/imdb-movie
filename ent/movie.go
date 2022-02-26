// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"imdb-movie/ent/movie"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Movie is the model entity for the Movie schema.
type Movie struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Rated holds the value of the "rated" field.
	Rated float32 `json:"rated,omitempty"`
	// RealeaseDate holds the value of the "realease_date" field.
	RealeaseDate time.Time `json:"realease_date,omitempty"`
	// Genre holds the value of the "genre" field.
	Genre string `json:"genre,omitempty"`
	// Poster holds the value of the "poster" field.
	Poster string `json:"poster,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Plot holds the value of the "plot" field.
	Plot string `json:"plot,omitempty"`
	// Stars holds the value of the "stars" field.
	Stars string `json:"stars,omitempty"`
	// ImdbRating holds the value of the "imdb_rating" field.
	ImdbRating string `json:"imdb_rating,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MovieQuery when eager-loading is set.
	Edges MovieEdges `json:"edges"`
}

// MovieEdges holds the relations/edges for other nodes in the graph.
type MovieEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e MovieEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Movie) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case movie.FieldRated:
			values[i] = new(sql.NullFloat64)
		case movie.FieldID:
			values[i] = new(sql.NullInt64)
		case movie.FieldTitle, movie.FieldGenre, movie.FieldPoster, movie.FieldDescription, movie.FieldPlot, movie.FieldStars, movie.FieldImdbRating:
			values[i] = new(sql.NullString)
		case movie.FieldRealeaseDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Movie", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Movie fields.
func (m *Movie) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case movie.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case movie.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				m.Title = value.String
			}
		case movie.FieldRated:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field rated", values[i])
			} else if value.Valid {
				m.Rated = float32(value.Float64)
			}
		case movie.FieldRealeaseDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field realease_date", values[i])
			} else if value.Valid {
				m.RealeaseDate = value.Time
			}
		case movie.FieldGenre:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field genre", values[i])
			} else if value.Valid {
				m.Genre = value.String
			}
		case movie.FieldPoster:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field poster", values[i])
			} else if value.Valid {
				m.Poster = value.String
			}
		case movie.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				m.Description = value.String
			}
		case movie.FieldPlot:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field plot", values[i])
			} else if value.Valid {
				m.Plot = value.String
			}
		case movie.FieldStars:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stars", values[i])
			} else if value.Valid {
				m.Stars = value.String
			}
		case movie.FieldImdbRating:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field imdb_rating", values[i])
			} else if value.Valid {
				m.ImdbRating = value.String
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Movie entity.
func (m *Movie) QueryUsers() *UserQuery {
	return (&MovieClient{config: m.config}).QueryUsers(m)
}

// Update returns a builder for updating this Movie.
// Note that you need to call Movie.Unwrap() before calling this method if this Movie
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Movie) Update() *MovieUpdateOne {
	return (&MovieClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Movie entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Movie) Unwrap() *Movie {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Movie is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Movie) String() string {
	var builder strings.Builder
	builder.WriteString("Movie(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", title=")
	builder.WriteString(m.Title)
	builder.WriteString(", rated=")
	builder.WriteString(fmt.Sprintf("%v", m.Rated))
	builder.WriteString(", realease_date=")
	builder.WriteString(m.RealeaseDate.Format(time.ANSIC))
	builder.WriteString(", genre=")
	builder.WriteString(m.Genre)
	builder.WriteString(", poster=")
	builder.WriteString(m.Poster)
	builder.WriteString(", description=")
	builder.WriteString(m.Description)
	builder.WriteString(", plot=")
	builder.WriteString(m.Plot)
	builder.WriteString(", stars=")
	builder.WriteString(m.Stars)
	builder.WriteString(", imdb_rating=")
	builder.WriteString(m.ImdbRating)
	builder.WriteByte(')')
	return builder.String()
}

// Movies is a parsable slice of Movie.
type Movies []*Movie

func (m Movies) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
