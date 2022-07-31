package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syafrin-ibrahim/bookstore_users_api/controllers/users"
)

func Mapurls() {
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ping pong")
	})

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.Createuser)
}
