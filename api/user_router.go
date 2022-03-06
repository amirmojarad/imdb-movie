package api

import (
	"fmt"
	"imdb-movie/crud"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddMoviesStruct struct {
	MovieIDs []int
	UserID   int
}

func (api API) DELETEUser(path string) {
	api.Router.DELETE(path, func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		res := api.Crud.DeleteUser(id)
		if res {
			ctx.IndentedJSON(http.StatusOK, fmt.Sprintf("Deleted user with id: %d", id))
		} else {
			ctx.IndentedJSON(http.StatusConflict, fmt.Sprintf("Can not delete user with id: %d", id))
		}
	})
}

func (api API) POSTUser(path string) {
	api.Router.POST(path, func(c *gin.Context) {
		/*
			{
				"email": "",
				"password": "",
				"fullName": "",
				"username": ""
			}
		*/
		userStruct := crud.UserStruct{}
		if err := c.BindJSON(&userStruct); err != nil {
			log.Println("On POSTUser Function api/user_router.go: ", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if createdUser, err := api.Crud.AddUser(&userStruct); err != nil {
			log.Println("On POSTUser Function api/user_router.go: ", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, createdUser)
		}
	})
}

func (api API) GETUser(path string) {
	api.Router.GET(path, func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("On GETUser Function in api/user_router.go: ", err)
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		user, err := api.Crud.GetUser(id)
		if err != nil {
			log.Println("On GETUser Function in api/user_router.go: ", err)
			c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("user with id %d not found", id))
			return
		}
		c.IndentedJSON(http.StatusOK, user)
	})
}

func (api API) GETUsers(path string) {
	api.Router.GET(path, func(c *gin.Context) {
		if allUsers, err := api.Crud.GetAllUsers(); err != nil {
			log.Println("On GETAllUsers in api/user_router.go: ", err)
			return
		} else {
			c.IndentedJSON(http.StatusOK, allUsers)
		}
	})
}

func (api API) POSTMoviesToUsers(path string) {
	api.Router.POST(path, func(c *gin.Context) {
		addMoviesStruct := AddMoviesStruct{}
		c.BindJSON(&addMoviesStruct)
		movies, err := api.Crud.AddMoviesToUser(&addMoviesStruct.MovieIDs, addMoviesStruct.UserID)
		if err != nil {
			log.Println("on POSTMoviesTo in api/user_router.go: ", err)
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, movies)
	})
}

func (api API) GETUsersMovies(path string) {
	api.Router.GET(path, func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		movies, err := api.Crud.GetUserMovie(id)
		if err != nil {
			log.Println("on GETUsersMovies in api/user_router.go: ", err)
			c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("user not found with id %d", id))
			return
		}
		c.IndentedJSON(http.StatusOK, movies)
	})
}

func (api API) GETUsersFavorites(path string) {
	api.Router.GET(path, func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		favorites, err := api.Crud.GetFavorites(id)
		if err != nil {
			log.Println("on GETUsersMovies in api/user_router.go: ", err)
			c.IndentedJSON(http.StatusNotFound, fmt.Sprintf("user not found with id %d", id))
			return
		}
		c.IndentedJSON(http.StatusOK, favorites)
	})
}

func (api API) POSTUsersFavorites(path string) {
	api.Router.POST(path, func(c *gin.Context) {
		addMoviesStruct := AddMoviesStruct{}
		c.BindJSON(&addMoviesStruct)
		favorites, err := api.Crud.AddFavorites(&addMoviesStruct.MovieIDs, addMoviesStruct.UserID)
		if err != nil {
			log.Println("on POSTUsersFavorites in api/user_router.go: ", err)
			c.IndentedJSON(http.StatusNotFound, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, favorites)
	})
}
