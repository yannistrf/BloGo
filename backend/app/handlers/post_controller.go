package handlers

import (
	"blogo/app/models"
	"blogo/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController interface {
	Add(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type postController struct {
	service services.PostService
}

func NewPostController(service services.PostService) PostController {
	return &postController{service: service}
}

func (controller *postController) Add(ctx *gin.Context) {
	var new_post models.Post
	err := ctx.ShouldBindJSON(&new_post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.service.Add(&new_post)
}

func (controller *postController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	post := controller.service.FindByID(uint(id))
	if post.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (controller *postController) FindAll(ctx *gin.Context) {
	posts := controller.service.FindAll()
	ctx.JSON(http.StatusOK, posts)
}

func (controller *postController) DeleteByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	controller.service.DeleteByID(uint(id))
}
