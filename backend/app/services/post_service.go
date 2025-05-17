package services

import (
	"blogo/app/models"
	"blogo/app/repositories"
)

type PostService interface {
	Add(*models.Post)
	FindByID(uint) *models.Post
	FindAll() *[]models.Post
	DeleteByID(uint)
}

type postService struct {
	repo repositories.PostRepo
}

func NewPostService(repo repositories.PostRepo) PostService {
	return &postService{repo: repo}
}

func (service *postService) Add(post *models.Post) {
	service.repo.Add(post)
}

func (service *postService) FindByID(id uint) *models.Post {
	return service.repo.FindByID(id)
}

func (service *postService) FindAll() *[]models.Post {
	return service.repo.FindAll()
}

func (service *postService) DeleteByID(id uint) {
	service.repo.DeleteByID(id)
}
