package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/database"
	"github.com/huzeyfebostan/myBlog/middlewares"
	"github.com/huzeyfebostan/myBlog/models"
)

func Admin(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := middlewares.ParseJwt(cookie)

	var user models.User

	database.DB().Where("id = ?", id).First(&user)

	return c.Render("admin", user)
}
