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
	StringSearch(ctx *gin.Context)
}

type postController struct {
	service services.PostService
}

func NewPostController(service services.PostService) PostController {
	return &postController{service: service}
}

func (controller *postController) Add(ctx *gin.Context) {
	var new_post models.Post
	new_post.UserID = ctx.GetUint("user_id")
	err := ctx.ShouldBindJSON(&new_post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.service.Add(&new_post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	page, err := strconv.Atoi(ctx.Query("page"))
	if page == 0 || err != nil {
		page = 1
	}
	posts := controller.service.FindAll(page)
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

func (controller *postController) StringSearch(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if page == 0 || err != nil {
		page = 1
	}

	posts := controller.service.StringSearch(ctx.Query("query"), page)
	ctx.JSON(http.StatusOK, posts)
}
