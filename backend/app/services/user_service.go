package services

import (
	"blogo/app/models"
	"blogo/app/repositories"
)

type UserService interface {
	Add(*models.User)
	FindByID(uint) *models.User
	FindAll() *[]models.User
	DeleteByID(uint)
}

type userService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userService{repo: repo}
}

func (service *userService) Add(user *models.User) {
	service.repo.Add(user)
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
