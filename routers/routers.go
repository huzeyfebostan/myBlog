package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/handlers"
	"github.com/huzeyfebostan/myBlog/middlewares"
)

func Setup(app *fiber.App) {

	app.Static("/", "./")

	//app.Get("/", handlers.MainPage)
	app.Get("/", handlers.AllArticle)
	//app.Get("/articles", handlers.AllArticle)

	//app.Get("/login", handlers.GetLogin)

	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.Register)

	//app.Get("/register", handlers.GetRegister)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/user", handlers.User)
	app.Get("/logout", handlers.Logout)

	app.Get("/users", handlers.AllUsers)
	app.Post("/users", handlers.CreateUser)
	app.Get("/users/:id", handlers.GetUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

	//app.Post("/update/:key", handlers.Update)
	//app.Get("/admin", handlers.User)
	//app.Get("/admin/:key", handlers.GetUser)

	app.Get("/article", handlers.GetArticle)
	app.Post("/article", handlers.CreateArticle)

	//app.Get("/adminUpdate/:key", handlers.GetUser)
	//app.Post("/adminUpdate/:key", handlers.Update)

	//app.Get("/update/:key", handlers.GetUser)

	app.Get("/delete/:key", handlers.Delete)
	app.Post("/update/:key", handlers.Delete)
}
