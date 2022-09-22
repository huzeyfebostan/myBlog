package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
)

type Article struct {
	gorm.Model
	WriterId uint   `json:"writerId"`
	ImgSrc   string `json:"imgSrc"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

type Entity interface {
	Count(db *gorm.DB) int64
	Take(db *gorm.DB, limit int, offset int) interface{}
}

func (article *Article) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Article{}).Count(&total)

	return total
}

func (article *Article) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []Article

	db.Offset(offset).Limit(limit).Find(&products)

	return products
}

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
	limit := 7
	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)

	total := entity.Count(db)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	}
}
