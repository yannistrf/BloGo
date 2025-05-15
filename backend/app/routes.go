package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutesInit(server *gin.Engine) {

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"health": "good"})
	})

}
