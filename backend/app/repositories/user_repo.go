package repositories

import (
	"blogo/app/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	Add(*models.User)
	FindByID(uint) models.User
	FindAll() []models.User
	DeleteByID(uint)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	db.AutoMigrate(models.User{})
	return &userRepo{db: db}
}

func (repo *userRepo) Add(user *models.User) {

}

func (repo *userRepo) FindByID(id uint) models.User {
	return models.User{}
}

func (repo *userRepo) FindAll() []models.User {
	return []models.User{}
}

func (repo *userRepo) DeleteByID(id uint) {

}
