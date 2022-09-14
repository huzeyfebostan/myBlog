package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/handlers"
)

func Setup(app *fiber.App) {

	app.Static("/", "./")

	app.Get("/", handlers.Mainpage)

	//app.Use(middlewares.IsAuthenticated)

	app.Get("/login", handlers.LoginGet)
	app.Post("/login", handlers.LoginPost)

	app.Get("/admin", handlers.Admin)
	//app.Post("/admin/:key", handlers.GetUser)
	app.Get("/update", handlers.Update)
	app.Get("/update/:key", handlers.GetUser)
	app.Post("/update/:key", handlers.GetUpdate)

	app.Get("/delete/:key", handlers.Delete)
	app.Post("/update/:key", handlers.Delete)

	app.Get("/success", handlers.Logout)

	app.Get("/unsuccess", handlers.Unsuccess)

	app.Get("/register", handlers.RegisterGet)
	app.Post("/register", handlers.RegisterPost)

}
