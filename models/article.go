package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	WriterId uint   `json:"writerId"`
	ImgSrc   string `json:"imgSrc"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Content  string `json:"content"`
	Category string `json:"category"`
}
