package handlers

import "github.com/gofiber/fiber/v2"

func Mainpage(c *fiber.Ctx) error {
	return c.Render("mainPage", fiber.Map{})
}
