package repositories

import (
	"blogo/app/models"

	"gorm.io/gorm"
)

type PostRepo interface {
	Add(*models.Post) error
	FindByID(uint) *models.Post
	FindAll() *[]models.Post
	DeleteByID(uint)
}

type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) PostRepo {
	return &postRepo{db: db}
}

func (repo *postRepo) Add(post *models.Post) error {
	var user models.User
	repo.db.First(&user, post.UserID)
	post.Author = user.Username
	return repo.db.Create(post).Error
}

func (repo *postRepo) FindByID(id uint) *models.Post {
	var post models.Post
	repo.db.First(&post, id)
	return &post
}

func (repo *postRepo) FindAll() *[]models.Post {
	var posts []models.Post
	repo.db.Find(&posts)
	return &posts
}

func (repo *postRepo) DeleteByID(id uint) {
	repo.db.Delete(&models.Post{}, id)
}
