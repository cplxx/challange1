package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title string `gorm:"title"`
	Color string `gorm:"color"`
}

var CategoryModel struct {
	Title string
	Color string
}
