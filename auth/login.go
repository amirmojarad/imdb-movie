package auth

import (
	"imdb-movie/auth/jwt"
	"imdb-movie/crud"
	"imdb-movie/dto"
	"imdb-movie/utils"
	"log"

	"github.com/gin-gonic/gin"
)

type LoginInformation struct {
	Email    string
	Password string
}

func LoginUser(email, password string, crud *crud.Crud) bool {
	fetchedUser, err := crud.GetUserByEmail(email)
	log.Println(fetchedUser)
	if err != nil {
		log.Println("on LoginUser in service/loginService.go: ", err)
		return false
	}
	b := utils.CheckPasswordHash(fetchedUser.Password, password)
	log.Println(b)
	if !b {
		log.Println("on LoginUser in service/loginService.go -- PASSWORD INCORRECT")
		return false
	}
	return true
}

func Login(ctx *gin.Context, crud *crud.Crud) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := LoginUser(credential.Email, credential.Password, crud)
	if isUserAuthenticated {
		return jwt.JWTAuthService().GenerateToken(credential.Email, true)
	}
	return ""
}
