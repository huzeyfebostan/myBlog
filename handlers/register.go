package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/database"
	"github.com/huzeyfebostan/myBlog/models"
)

func GetRegister(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}

func Register(c *fiber.Ctx) error {

	var request models.RequestRegister
	var user models.User

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	/*if err := database.DB().Where("email = ?", request.Email).First(&user).Error; err == nil {
		return errors.New("This mail has been used")
	}*/

	if request.Password != request.PasswordConfirm {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user = models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}

	user.SetPassword(user.Password)
	user.RoleId = 2

	database.DB().Create(&user)

	/*err := database.DB().Create(&user).Error
	if err != nil {
		return errors.New("unable to create user")
	}*/

	return c.JSON(user)
}
