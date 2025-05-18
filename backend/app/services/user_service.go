package services

import (
	"blogo/app/models"
	"blogo/app/repositories"
)

type UserService interface {
	Add(*models.User) error
	FindByID(uint) *models.User
	FindAll() *[]models.User
	DeleteByID(uint)
	FindPostsByID(uint) *[]models.Post
}

type userService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userService{repo: repo}
}

func (service *userService) Add(user *models.User) error {
	return service.repo.Add(user)
}

func (service *userService) FindByID(id uint) *models.User {
	return service.repo.FindByID(id)
}

func (service *userService) FindAll() *[]models.User {
	return service.repo.FindAll()
}

func (service *userService) DeleteByID(id uint) {
	service.repo.DeleteByID(id)
}

func (service *userService) FindPostsByID(id uint) *[]models.Post {
	return service.repo.FindPostsByID(id)
}
