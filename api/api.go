package api

import (
	"context"
	"imdb-movie/crud"
	"imdb-movie/ent"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router *gin.Engine
	Crud   *crud.Crud
}

func RunAPI(ctx context.Context, client *ent.Client) {
	api := API{Router: gin.Default(), Crud: &crud.Crud{Ctx: ctx, Client: client}}

	// User Router
	api.POSTUser("/users")
	api.GETUser("/users")
	api.GETUsers("/users/:id")
	api.POSTMoviesToUsers("/user/movies")
	api.GETUsersMovies("/user/:id/movies")
	// Movie Router
	api.POSTMovies("/movies")
	api.GETMovies("/movies")
	api.GETMovieByID("/movies/:id")
	api.Router.Run("localhost:8080")
}
