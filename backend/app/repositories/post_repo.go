package repositories

import (
	"blogo/app/models"

	"gorm.io/gorm"
)

const PAGE_SIZE = 3

type PostRepo interface {
	Add(*models.Post) error
	FindByID(uint) *models.Post
	FindAll(int) *[]models.Post
	DeleteByID(uint)
	StringSearch(string, int) *[]models.Post
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

func (repo *postRepo) FindAll(page int) *[]models.Post {
	var posts []models.Post
	offset := (page - 1) * PAGE_SIZE
	repo.db.Order("created_at DESC").Offset(offset).Limit(PAGE_SIZE).Find(&posts)
	return &posts
}

func (repo *postRepo) DeleteByID(id uint) {
	repo.db.Delete(&models.Post{}, id)
}

func (repo *postRepo) StringSearch(query string, page int) *[]models.Post {
	query = "%" + query + "%" // add the wildcards
	var posts []models.Post
	offset := (page - 1) * PAGE_SIZE
	repo.db.Where("title LIKE ? OR content LIKE ?", query, query).
		Order("created_at DESC").Offset(offset).Limit(PAGE_SIZE).Find(&posts)
	return &posts
}
