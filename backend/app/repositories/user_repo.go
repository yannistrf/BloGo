package repositories

import (
	"blogo/app/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	Add(*models.User)
	FindByID(uint) *models.User
	FindAll() *[]models.User
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
	repo.db.Create(user)
}

func (repo *userRepo) FindByID(id uint) *models.User {
	var user models.User
	repo.db.First(&user, id)
	return &user
}

func (repo *userRepo) FindAll() *[]models.User {
	var users []models.User
	repo.db.Find(&users)
	return &users
}

func (repo *userRepo) DeleteByID(id uint) {
	repo.db.Delete(&models.User{}, id)
}
