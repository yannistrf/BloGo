package main

import (
	"blogo/app"
	"blogo/app/repositories"
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
	userRepo.DeleteByID(3)

	server := gin.Default()
	app.RoutesInit(server)
	server.Run(":8081")
}
