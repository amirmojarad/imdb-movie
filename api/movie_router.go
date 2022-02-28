package api

import (
	"imdb-movie/api/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostMovieStruct struct {
	Title string
}

func (api API) GETMovieByID(path string) {

	api.Router.GET(path, func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		movie, err := api.Crud.Client.Movie.Get(api.Crud.Ctx, id)
		if err != nil {
			log.Println("on GETMovie Function in api/movie_router.go: ", err)
			return
		}
		c.IndentedJSON(http.StatusOK, movie)
	})
}

func (api API) GETMovies(path string) {
	api.Router.GET(path, func(c *gin.Context) {
		movies, err := api.Crud.GetAllMovies()
		if err != nil {
			log.Println("on GETMovies Function in api/movie_router.go: ", err)
			return
		}
		c.IndentedJSON(http.StatusOK, movies)
	})
}

func (api API) POSTMovies(path string) {
	api.Router.POST(path, func(c *gin.Context) {
		var postMovieStruct PostMovieStruct
		c.BindJSON(&postMovieStruct)
		movies, err := utils.GETMovies(postMovieStruct.Title)
		if err != nil {
			log.Println("on POSTMovies Function in api/movie_router.go: ", err)
			return
		}
		fetchedMovies, err := api.Crud.AddAllMovies(movies)
		if err != nil {
			log.Println("on POSTMovies Function in api/movie_router.go: ", err)
			return
		}
		c.IndentedJSON(http.StatusOK, fetchedMovies)
	})
}
