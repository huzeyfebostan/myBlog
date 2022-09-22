package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func MainPage(c *fiber.Ctx) error {
	return c.Render("mainPage", fiber.Map{})
}
