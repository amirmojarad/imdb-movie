package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().MinLen(4),
		field.String("email").Unique(),
		field.String("full_name").Optional().Nillable(),
		field.String("password").Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("movies", Movie.Type),
		edge.To("favorites", Movie.Type),
	}
}
