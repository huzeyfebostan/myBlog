package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Id       uint   `json:"id"`
	WriterId uint   `json:"writerId"`
	ImgSrc   string `json:"imgSrc"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Head     string `json:"head"`
	View     uint   `json:"view"`
	Category string `json:"category"`
}
