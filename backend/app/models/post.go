package models

type Post struct {
	ID      uint   `json:"id" gorm:"primary_key;auto_increment"`
	Title   string `json:"title" binding:"required" gorm:"type:varchar(64)"`
	Content string `json:"content" binding:"required" gorm:"type:varchar(256)"`
}
