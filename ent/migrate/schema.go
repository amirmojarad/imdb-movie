// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MoviesColumns holds the columns for the "movies" table.
	MoviesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "rated", Type: field.TypeFloat32},
		{Name: "realease_date", Type: field.TypeTime},
		{Name: "genre", Type: field.TypeString},
		{Name: "poster", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "plot", Type: field.TypeString},
		{Name: "stars", Type: field.TypeString},
		{Name: "imdb_rating", Type: field.TypeString},
		{Name: "user_movies", Type: field.TypeInt, Nullable: true},
		{Name: "user_favorites", Type: field.TypeInt, Nullable: true},
	}
	// MoviesTable holds the schema information for the "movies" table.
	MoviesTable = &schema.Table{
		Name:       "movies",
		Columns:    MoviesColumns,
		PrimaryKey: []*schema.Column{MoviesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "movies_users_movies",
				Columns:    []*schema.Column{MoviesColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "movies_users_favorites",
				Columns:    []*schema.Column{MoviesColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "full_name", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MoviesTable,
		UsersTable,
	}
)

func init() {
	MoviesTable.ForeignKeys[0].RefTable = UsersTable
	MoviesTable.ForeignKeys[1].RefTable = UsersTable
}
