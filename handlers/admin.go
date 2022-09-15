package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/database"
	"github.com/huzeyfebostan/myBlog/middlewares"
	"github.com/huzeyfebostan/myBlog/models"
	"strconv"
)

func Admin(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := middlewares.ParseJwt(cookie)

	var user models.User

	database.DB().Where("id = ?", id).First(&user)

	return c.Render("admin", user)
}

func AdminUpdate(c *fiber.Ctx) error {
	/*if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	id, _ := middlewares.ParseJwt(cookie)

	var user models.User

	database.DB().Where("id = ?", id).First(&user)

	return c.Render("adminUpdate", user)*/
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var request models.RequestRegister

	id, _ := strconv.Atoi(c.Params("key"))

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if request.Password != request.PasswordConfirm {
		return errors.New("passwords do not match")
	}

	user := models.User{
		Id:        uint(id),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}

	if request.Password != "" {
		user.SetPassword(user.Password)
	}

	database.DB().Model(&user).Where("id = ?", user.Id).Updates(user)

	return c.Redirect("/admin")
}
