package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title      string `json:"title"`
	Content    string `json:"content"`
	Slug       string `json:"slug"`
	CategoryId uint   `json:"category_id"`
}

func init() {
	DB.AutoMigrate(&Post{})
}
