package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/handlers"
	"github.com/huzeyfebostan/myBlog/middlewares"
)

func Setup(app *fiber.App) {

	app.Static("/", "./")

	app.Get("/", handlers.Mainpage)

	app.Get("/login", handlers.GetLogin)
	app.Post("/login", handlers.PostLogin)

	app.Get("/register", handlers.GetRegister)
	app.Post("/register", handlers.PostRegister)

	app.Use(middlewares.IsAuthenticated)
	//TODO: Şifre gizlenecek (sor)
	//TODO: Aynı kullanıcı kayıt kontrolu yap
	app.Get("/user", handlers.User)
	app.Get("/user/:key", handlers.GetUser)

	app.Get("/admin", handlers.Admin)

	app.Get("/adminUpdate", handlers.AdminUpdate)
	//app.Get("adminUpdate/:key", handlers.GetUser)
	app.Get("/update", handlers.GetUpdate)
	app.Get("/update/:key", handlers.GetUser)
	app.Post("/update/:key", handlers.Update)

	app.Get("/delete/:key", handlers.Delete)
	app.Post("/update/:key", handlers.Delete)

	app.Get("/success", handlers.Logout)

	app.Get("/unsuccess", handlers.Unsuccess)
}
