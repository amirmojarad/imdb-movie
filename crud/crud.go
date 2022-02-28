package crud

import (
	"context"
	"imdb-movie/ent"
)

type Crud struct {
	Ctx    context.Context
	Client *ent.Client
}
