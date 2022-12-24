package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/handlers"
	"github.com/huzeyfebostan/myBlog/middlewares"
)

func Setup(app *fiber.App) {

	//app.Static("/", "./")

	app.Post("/api/login", handlers.Login)
	app.Post("/api/register", handlers.Register)

	app.Get("/api/articles", handlers.AllArticles)
	app.Get("/api/article/:id", handlers.GetArticle)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", handlers.User)
	app.Get("/api/logout", handlers.Logout)

	app.Get("/api/users", handlers.AllUsers)
	app.Post("/api/users", handlers.CreateUser)
	app.Get("/api/users/:id", handlers.GetUser)
	app.Put("/api/users/:id", handlers.UpdateUser)
	app.Delete("/api/users/:id", handlers.DeleteUser)

	app.Post("/api/articles", handlers.CreateArticle)
	app.Get("/api/articles/:id", handlers.ActiveUserArticles)
	app.Put("/api/articles/:id", handlers.UpdateArticle)
	app.Delete("/api/articles/:id", handlers.DeleteArticle)
}
