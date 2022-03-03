package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Float32("rated"),
		field.Time("realease_date").Default(time.Now),
		field.String("genre"),
		field.String("poster"),
		field.String("description"),
		field.String("plot"),
		field.String("stars"),
		field.String("imdb_rating"),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return []ent.Edge{}
}
