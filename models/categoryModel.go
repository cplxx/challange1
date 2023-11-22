package models

type Category struct {
	ID    uint   `gorm:"primaryKey"`
	title string `gorm:"title"`
	color string `gorm:"color"`
}
