package main

import (
	"blogo/app"
	"blogo/app/handlers"
	"blogo/app/models"
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
	db.Exec("PRAGMA foreign_keys = ON") // for sqlite
	db.AutoMigrate(models.User{}, models.Post{})

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := handlers.NewUserController(userService)

	postRepo := repositories.NewPostRepo(db)
	postService := services.NewPostService(postRepo)
	postController := handlers.NewPostController(postService)

	authController := handlers.NewAuthController(userService)

	server := gin.Default()
	app.RoutesInit(server, userController, postController, authController)
	server.Run(":8081")
}
