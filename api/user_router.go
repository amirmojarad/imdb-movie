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

// import (
// 	"imdb-movie/ent"
// 	"imdb-movie/ent/user"
// 	"log"
// )

// func (userRouter *API) GetUser(id int) (*ent.User, error) {
// 	fetchedUser, err := userRouter.Client.User.Get(userRouter.Ctx, id)
// 	if err != nil {
// 		log.Println("error when fetching user by id: ", err)
// 		return nil, err
// 	} else if fetchedUser == nil {
// 		log.Println("no user with id: ", id)
// 		return nil, nil
// 	}
// 	return fetchedUser, nil
// }

// func (userRouter *API) PostUsers() (*ent.User, error) {
// 	// newUser, err := userRouter.
// 	// 			Client.User.Create().
// 	// 			SetUsername().
// 	// 			Save(userRouter.Ctx)

// 	return nil, nil
// }

// func (userRouter *API) GetUserMovies(id int) ([]*ent.Movie, error) {
// 	user, err := userRouter.Client.User.Get(userRouter.Ctx, id)
// 	log.Println(user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	movies, err := user.QueryMovies().All(userRouter.Ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return movies, nil
// }

// func (userRouter *API) AddMovies2User(user *ent.User, moviesIDs []int) {
// 	for _, id := range moviesIDs {
// 		userRouter.Client.User.UpdateOne(user).AddMovieIDs(id).Save(userRouter.Ctx)
// 	}
// }

// func (userRouter *API) GetUsers() []*ent.User {
// 	users, err := userRouter.Client.User.Query().Order(ent.Asc(user.FieldID)).All(userRouter.Ctx)
// 	if err != nil {
// 		log.Println("on GetUsers error occured: ", err)
// 	}
// 	return users
// }
