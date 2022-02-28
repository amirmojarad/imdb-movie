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
	api.POSTUser("/users")
	api.GETUser("/users/:id")
	api.GETUsers("/users")
	api.POSTMoviesToUsers("/users/movies")
	api.Router.Run("localhost:8080")
}

// func (api API) RunAPI() {
// 	router := gin.Default()

// 	router.GET("/users/movies/:id", func(c *gin.Context) {
// 		userID, _ := strconv.Atoi(c.Param("id"))
// 		movies, err := api.GetUserMovies(userID)
// 		log.Println(movies)
// 		if err != nil {
// 			log.Println("on get user movies: ", err)
// 			return
// 		}
// 		c.IndentedJSON(http.StatusOK, movies)

// 	})

// 	router.POST("/users/movies", func(c *gin.Context) {
// 		/*
// 			{
// 				"movieIDs": [],
// 				"userID": 1
// 			}
// 		*/
// 		addMovies := AddMoviesStruct{}
// 		if err := c.BindJSON(&addMovies); err != nil {
// 			log.Println("on Add Movies endpoint: ", err)
// 			c.IndentedJSON(http.StatusBadRequest, nil)
// 		}
// 		log.Println(addMovies)
// 		user, err := api.GetUser(addMovies.UserID)
// 		if err != nil {
// 			log.Println("cannot find user on AdMovies endpoint: ", err)
// 			c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("cannot find user with %d", addMovies.UserID)})
// 		}
// 		api.AddMovies2User(user, addMovies.MovieIDs)
// 		c.IndentedJSON(http.StatusOK, "its ok!")
// 	})

// 	router.GET("/movies/:id", func(c *gin.Context) {
// 		id, _ := strconv.Atoi(c.Param("id"))
// 		if movie, err := api.GetMovie(id); err != nil {
// 			log.Println(err)
// 			return
// 		} else {
// 			c.IndentedJSON(http.StatusOK, movie)
// 		}
// 	})

// 	router.GET("/movies/imdb/", func(c *gin.Context) {
// 		response := GetMovie("batman")
// 		if movies, err := api.AddMovies(&response); err != nil {
// 			log.Println(err)
// 		} else {
// 			log.Println("MVOEIS: ", movies)
// 		}
// 		c.IndentedJSON(http.StatusOK, response)
// 	})

// 	router.GET("/movies/all", func(c *gin.Context) {
// 		movies, err := api.GetAllMovies()
// 		if err != nil {
// 			return
// 		}
// 		c.IndentedJSON(http.StatusOK, movies)
// 	})

// 	router.GET("/users", func(c *gin.Context) {
// 		users := api.GetUsers()
// 		c.IndentedJSON(http.StatusOK, users)
// 	})

// 	router.GET("/users/:id", func(c *gin.Context) {
// 		id, _ := strconv.Atoi(c.Param("id"))
// 		user, err := api.GetUser(id)
// 		if err != nil {
// 			log.Println("error when fetching user with id: ", id)
// 			c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("no user found with %v", id)})
// 			return
// 		}
// 		log.Println("user with id", id, "founded: ", user)
// 		c.IndentedJSON(http.StatusOK, user)
// 	})

// 	router.POST("/users", func(c *gin.Context) {
// 		// TODO : separate this lines of code into user_router.go
// 		temp := UserStruct{}

// 		if err := c.BindJSON(&temp); err != nil {
// 			log.Println("error!!!!!", err)
// 		}
// 		log.Println(temp)
// 		u, err := api.Client.User.Create().
// 			SetEmail(temp.Email).SetFullName(temp.FullName).SetPassword(temp.Password).SetUsername(temp.Username).Save(api.Ctx)
// 		if err != nil {
// 			log.Println("error in post new user method: ", err)
// 			c.IndentedJSON(http.StatusOK, gin.H{"message": "an error occured", "detail": err.Error()})
// 		} else {
// 			c.IndentedJSON(http.StatusOK, u)
// 		}
// 	})

// 	router.Run("localhost:8080")

// }
