package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;auto_increment"`
	Username string `json:"username" binding:"required" gorm:"type:varchar(32);unique"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(32)"`
	Posts    []Post `json:"posts" gorm:"constraint:OnDelete:CASCADE;"`
}
