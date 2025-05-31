package main

import (
	"blogo/app"
	"blogo/app/handlers"
	"blogo/app/models"
	"blogo/app/repositories"
	"blogo/app/services"
	"blogo/app/utils"
	"errors"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	var db_exists bool
	if _, err := os.Stat("blogo.db"); errors.Is(err, os.ErrNotExist) {
		db_exists = false
	} else {
		db_exists = true
	}

	db, err := gorm.Open(sqlite.Open("blogo.db"), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	db.Exec("PRAGMA foreign_keys = ON") // for sqlite
	db.AutoMigrate(models.User{}, models.Post{}, models.Comment{})

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := handlers.NewUserController(userService)

	postRepo := repositories.NewPostRepo(db)
	postService := services.NewPostService(postRepo)
	postController := handlers.NewPostController(postService)

	authController := handlers.NewAuthController(userService)

	if !db_exists {
		utils.InsertTestData(userRepo, postRepo)
	}

	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // TODO: read from env
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	app.RoutesInit(server, userController, postController, authController)
	server.Run("127.0.0.1:8081")
}
