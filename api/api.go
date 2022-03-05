package api

import (
	"context"
	"imdb-movie/controller"
	"imdb-movie/crud"
	"imdb-movie/ent"
	"imdb-movie/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router *gin.Engine
	Crud   *crud.Crud
}

func (api API) Options(path string) {
	api.Router.OPTIONS(path, func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return

		}
		c.Next()
	})

}

func RunAPI(ctx context.Context, client *ent.Client) {

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	api := API{Router: gin.New(), Crud: &crud.Crud{Ctx: ctx, Client: client}}

	api.Router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// User Router
	api.POSTUser("/users")
	api.GETUser("/users")
	api.GETUsers("/users/:id")
	api.POSTMoviesToUsers("/user/movies")
	api.GETUsersMovies("/user/:id/movies")
	api.GETUsersFavorites("/user/:id/favorites")
	api.POSTUsersFavorites("/user/:id/favorites")
	// Movie Router
	api.POSTMovies("/movies")
	api.GETMovies("/movies")
	api.GETMovieByID("/movies/:id")
	// Options
	api.Options("/options")
	api.Router.Run("localhost:8080")
}
