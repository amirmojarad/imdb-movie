package api

import (
	"imdb-movie/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api API) Login(path string) {
	api.Router.POST(path, func(ctx *gin.Context) {
		token := auth.Login(ctx, api.Crud)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
}
