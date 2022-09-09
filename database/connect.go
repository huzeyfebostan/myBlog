package database

import (
	"github.com/huzeyfebostan/myBlog/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func Connect() {
	dns := "host=localhost user=hoze password=123 dbname=myblogdb port=5432 sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	db.AutoMigrate(&models.User{})
}
