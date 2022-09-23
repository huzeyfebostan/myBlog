package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/database"
	"github.com/huzeyfebostan/myBlog/middlewares"
	"github.com/huzeyfebostan/myBlog/models"
)

func GetArticle(c *fiber.Ctx) error {
	return c.Render("article", fiber.Map{})
}

func CreateArticle(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var requestArticle models.RequestArticle

	if err := c.BodyParser(&requestArticle); err != nil {
		return err
	}

	article := models.Article{
		ImgSrc:   requestArticle.ImgSrc,
		Title:    requestArticle.Title,
		Summary:  requestArticle.Summary,
		Content:  requestArticle.Content,
		Category: requestArticle.Category,
	}

	if requestArticle.ImgSrc == "" {
		requestArticle.ImgSrc = "../assets/image/article.png"
	}

	database.DB().Create(&article)

	return c.Redirect("/")
}

func AllArticle(c *fiber.Ctx) error {
	var Article []models.Article

	database.DB().Find(&Article)

	return c.Render("mainPage", Article)
}
