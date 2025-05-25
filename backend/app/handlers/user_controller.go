package handlers

import (
	"blogo/app/models"
	"blogo/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Add(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
	FindPostsByID(ctx *gin.Context)
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{service: service}
}

func (controller *userController) Add(ctx *gin.Context) {
	var new_user models.User
	err := ctx.ShouldBindJSON(&new_user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.service.Add(&new_user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (controller *userController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user := controller.service.FindByID(uint(id))
	if user.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (controller *userController) FindAll(ctx *gin.Context) {
	users := controller.service.FindAll()
	ctx.JSON(http.StatusOK, users)
}

func (controller *userController) DeleteByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	controller.service.DeleteByID(uint(id))
}

func (controller *userController) FindPostsByID(ctx *gin.Context) {
	posts := controller.service.FindPostsByID(ctx.GetUint("user_id"))
	ctx.JSON(http.StatusOK, posts)
}
