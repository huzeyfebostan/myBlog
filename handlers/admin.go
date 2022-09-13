package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Admin(c *fiber.Ctx) error {
	/*var usr AdminPage

	user, _ := Find()

	usr = AdminPage{
		User: user,
	}*/
	return c.Render("admin", fiber.Map{})
}
