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
		field.String("year").Default("1900"),
		field.Float32("rated"),
		field.Time("realease_date").Default(time.Now),
		field.String("genre"),
		field.String("language"),
		field.String("poster"),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return nil
}
