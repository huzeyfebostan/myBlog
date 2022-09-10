package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/database"
	"github.com/huzeyfebostan/myBlog/models"
)

func RegisterGet(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}

func RegisterPost(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	//TODO: Aynı kullanıcı kayıt kontrolu yap
	user.SetPassword(user.Password)

	err := database.DB().Create(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return c.Redirect("/login")
}
