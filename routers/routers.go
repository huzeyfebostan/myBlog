package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/handlers"
)

func Setup(app *fiber.App) {

	app.Static("/", "./")

	app.Get("/", handlers.Mainpage)

	app.Get("/login", handlers.LoginGet)
	app.Post("/login", handlers.LoginPost)

	app.Get("/success", handlers.Logout)

	app.Get("/admin", handlers.Admin)

	app.Get("/unsuccess", handlers.Unsuccess)

	app.Get("/register", handlers.RegisterGet)
	app.Post("/register", handlers.RegisterPost)

}
