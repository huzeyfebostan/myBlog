package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/database"
	"github.com/huzeyfebostan/myBlog/middlewares"
	"github.com/huzeyfebostan/myBlog/models"
	"gorm.io/gorm"
	"strconv"
)

func CreateArticle(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}
	var article models.Article

	if err := c.BodyParser(&article); err != nil {
		return err
	}

	article.WriterId = middlewares.GetUserId(c)

	database.DB().Create(&article)

	return c.JSON(fiber.Map{
		"message": "Makale başarıyla oluşturuldu",
	})
}

func AllArticles(c *fiber.Ctx) error {
	var articles []models.Article

	database.DB().Find(&articles)

	return c.JSON(articles)
}

func ActiveUserArticles(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}
	var LastLog uint
	log := models.UserLog{}

	id, _ := strconv.Atoi(c.Params("id"))
	database.DB().Select("UserId").Last(&log).Scan(&LastLog)

	article := models.Article{
		WriterId: uint(id),
	}
	database.DB().Where("writer_id = ?", article.WriterId).Find(&article)
	if article.WriterId != LastLog {
		return c.JSON(fiber.Map{
			"message": "Yetkiniz yok",
		})
	}

	return c.JSON(article)
}

func GetArticle(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	article := models.Article{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	database.DB().Preload("ID").Find(&article)
	return c.JSON(article)
}

func UpdateArticle(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	var reqArticle models.RequestArticle

	if err := c.BodyParser(&reqArticle); err != nil {
		return err
	}

	article := models.Article{
		Model:    gorm.Model{ID: uint(id)},
		ImgSrc:   reqArticle.ImgSrc,
		Title:    reqArticle.Title,
		Summary:  reqArticle.Summary,
		Content:  reqArticle.Content,
		Category: reqArticle.Category,
	}

	database.DB().Model(&article).Where("id = ?", article.ID).Updates(article)

	return c.JSON(article)
}

func DeleteArticle(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	article := models.Article{
		Model: gorm.Model{ID: uint(id)},
	}

	database.DB().Delete(&article)

	return c.JSON(fiber.Map{
		"message": "Makale başarıyla silindi",
	})
}

//func Article(c *fiber.Ctx) error {}
