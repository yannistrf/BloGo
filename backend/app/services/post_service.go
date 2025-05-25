package services

import (
	"blogo/app/models"
	"blogo/app/repositories"
)

type PostService interface {
	Add(*models.Post) error
	FindByID(uint) *models.Post
	FindAll() *[]models.Post
	DeleteByID(uint)
	StringSearch(string) *[]models.Post
}

type postService struct {
	repo repositories.PostRepo
}

func NewPostService(repo repositories.PostRepo) PostService {
	return &postService{repo: repo}
}

func (service *postService) Add(post *models.Post) error {
	return service.repo.Add(post)
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

func (service *postService) StringSearch(query string) *[]models.Post {
	return service.repo.StringSearch(query)
}
