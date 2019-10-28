package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

func init() {
	DB.AutoMigrate(&Category{})
}
