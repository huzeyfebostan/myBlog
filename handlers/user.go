package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/database"
	"github.com/huzeyfebostan/myBlog/middlewares"
	"github.com/huzeyfebostan/myBlog/models"
	"strconv"
)

func CreateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("123")
	user.RoleId = 2

	database.DB().Create(&user)

	return c.JSON(&user)
}

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB().Find(&users)

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB().Preload("id").Find(&user)
	if user.RoleId == 0 {
		return c.JSON(fiber.Map{
			"message": "Böyle bir kullanıcı yok",
		})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var request models.RequestRegister

	id, _ := strconv.Atoi(c.Params("id"))

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if request.Password != request.PasswordConfirm {
		return c.JSON(fiber.Map{
			"message": "Şifreler uyuşmuyor aloo",
		})
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

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB().Delete(&user)

	return nil
}
