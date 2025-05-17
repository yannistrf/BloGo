package main

import (
	"blogo/app"
	"blogo/app/handlers"
	"blogo/app/repositories"
	"blogo/app/services"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := gorm.Open(sqlite.Open("blogo.db"), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := handlers.NewUserController(userService)

	postRepo := repositories.NewPostRepo(db)
	postService := services.NewPostService(postRepo)
	postController := handlers.NewPostController(postService)

	server := gin.Default()
	app.RoutesInit(server, userController, postController)
	server.Run(":8081")
}
