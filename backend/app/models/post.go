package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment"`
	Title     string    `json:"title" binding:"required" gorm:"type:varchar(64)"`
	Content   string    `json:"content" binding:"required" gorm:"type:varchar(256)"`
	CreatedAt time.Time `json:"created_at"` // gorm fills this field by default with NOW value
	Author    string    `json:"author" gorm:"type:varchar(32)"`
	UserID    uint      `json:"user_id" gorm:"foreignKey"`
	Comments  []Comment `json:"comments"`
}

type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content" binding:"required" gorm:"type:varchar(256)"`
	Author    string    `json:"author" gorm:"type:varchar(32)"`
	CreatedAt time.Time `json:"created_at"` // gorm fills this field by default with NOW value
}
