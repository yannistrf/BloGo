package handlers

import (
	"blogo/app/models"
	"blogo/app/services"
	"blogo/app/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	userService services.UserService
}

func NewAuthController(service services.UserService) AuthController {
	return &authController{userService: service}
}

func (controller *authController) Login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result_user := controller.userService.FindByUsername(user.Username)
	if result_user.ID == 0 || result_user.Password != user.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad credentials"})
		return
	}

	tokenString, err := utils.GenerateJWT(result_user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (controller *authController) Register(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.userService.Add(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
