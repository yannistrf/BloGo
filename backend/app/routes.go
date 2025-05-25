package app

import (
	"blogo/app/handlers"
	"blogo/app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutesInit(server *gin.Engine,
	userController handlers.UserController,
	postController handlers.PostController,
	authController handlers.AuthController) {

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"health": "good"})
	})

	post_routes := server.Group("/post", middlewares.AuthorizeJWT)
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

	user_routes := server.Group("/user", middlewares.AuthorizeJWT)
	{
		user_routes.POST("/add", func(ctx *gin.Context) {
			userController.Add(ctx)
		})

		user_routes.GET("/:id", func(ctx *gin.Context) {
			userController.FindByID(ctx)
		})

		user_routes.GET("/all", func(ctx *gin.Context) {
			userController.FindAll(ctx)
		})

		user_routes.DELETE("/:id", func(ctx *gin.Context) {
			userController.DeleteByID(ctx)
		})

		user_routes.GET("/me", func(ctx *gin.Context) {
			userController.FindPostsByID(ctx)
		})
	}

	auth_routes := server.Group("/auth")
	{
		auth_routes.POST("/login", func(ctx *gin.Context) {
			authController.Login(ctx)
		})

		auth_routes.POST("/register", func(ctx *gin.Context) {
			authController.Register(ctx)
		})
	}
}
