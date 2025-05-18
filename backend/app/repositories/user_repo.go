package repositories

import (
	"blogo/app/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	Add(*models.User) error
	FindByID(uint) *models.User
	FindAll() *[]models.User
	DeleteByID(uint)
	FindPostsByID(uint) *[]models.Post
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (repo *userRepo) Add(user *models.User) error {
	return repo.db.Create(user).Error
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

func (repo *userRepo) FindPostsByID(id uint) *[]models.Post {
	var posts []models.Post
	repo.db.Where(&models.Post{UserID: id}).Find(&posts)
	return &posts
}
