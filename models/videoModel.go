package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	Url         string `gorm:"url"`
}

var VideoModel struct {
	Title       string
	Description string
	Url         string
}
