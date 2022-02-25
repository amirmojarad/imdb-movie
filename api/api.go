package api

import (
	"context"
	"fmt"
	"imdb-movie/ent"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type API struct {
	Ctx    context.Context
	Client *ent.Client
}

type UserStruct struct {
	Email    string
	Password string
	FullName string
	Username string
}

func (api API) RunAPI() {
	router := gin.Default()
	router.GET("/users", func(c *gin.Context) {
		users := api.GetUsers()
		c.IndentedJSON(http.StatusOK, users)
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		user, err := api.GetUser(id)
		if err != nil {
			log.Println("error when fetching user with id: ", id)
			c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("no user found with %v", id)})
			return
		}
		log.Println("user with id", id, "founded: ", user)
		c.IndentedJSON(http.StatusOK, user)
	})

	router.POST("/users", func(c *gin.Context) {

		temp := UserStruct{}

		if err := c.BindJSON(&temp); err != nil {
			log.Println("error!!!!!", err)
		}
		log.Println(temp)
		u, err := api.Client.User.Create().
			SetEmail(temp.Email).SetFullName(temp.FullName).SetPassword(temp.Password).SetUsername(temp.Username).Save(api.Ctx)
		if err != nil {
			log.Println("error in post new user method: ", err)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "an error occured", "detail": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, u)
		}
	})

	router.Run("localhost:8080")

}
