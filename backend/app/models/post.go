package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment"`
	Title     string    `json:"title" binding:"required" gorm:"type:varchar(64)"`
	Content   string    `json:"content" binding:"required" gorm:"type:varchar(256)"`
	CreatedAt time.Time `json:"created_at"` // gorm fills this field by default with NOW value
	UserID    uint      `json:"user_id" binding:"required" gorm:"foreignKey"`
}
