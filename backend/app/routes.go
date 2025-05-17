package app

import (
	"blogo/app/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutesInit(server *gin.Engine, postController handlers.PostController) {
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"health": "good"})
	})

	post_routes := server.Group("/post")
	{
		post_routes.POST("/add", func(ctx *gin.Context) {
			postController.Add(ctx)
		})

		post_routes.GET("/:id", func(ctx *gin.Context) {
			postController.FindByID(ctx)
		})

		post_routes.GET("/all", func(ctx *gin.Context) {
			postController.FindAll(ctx)
		})

		post_routes.DELETE("/:id", func(ctx *gin.Context) {
			postController.DeleteByID(ctx)
		})
	}
}
