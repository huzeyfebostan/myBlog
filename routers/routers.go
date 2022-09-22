package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/handlers"
	"github.com/huzeyfebostan/myBlog/middlewares"
)

func Setup(app *fiber.App) {

	app.Static("/", "./")

	app.Get("/", handlers.MainPage)

	app.Get("/login", handlers.GetLogin)
	app.Post("/login", handlers.PostLogin)

	app.Get("/register", handlers.GetRegister)
	app.Post("/register", handlers.PostRegister)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/user", handlers.User)
	app.Get("/user/:key", handlers.GetUser)

	app.Get("/admin", handlers.User)
	app.Get("/admin/:key", handlers.GetUser)

	app.Get("/article", handlers.GetArticle)
	app.Post("/article", handlers.CreateArticle)

	app.Get("/adminUpdate/:key", handlers.GetUser)
	app.Post("/adminUpdate/:key", handlers.Update)

	app.Get("/update/:key", handlers.GetUser)
	app.Post("/update/:key", handlers.Update)

	app.Get("/delete/:key", handlers.Delete)
	app.Post("/update/:key", handlers.Delete)

	app.Get("/", handlers.Logout)
}
