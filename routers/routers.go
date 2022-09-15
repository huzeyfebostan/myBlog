package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/handlers"
	"github.com/huzeyfebostan/myBlog/middlewares"
)

func Setup(app *fiber.App) {

	app.Static("/", "./")

	app.Get("/", handlers.Mainpage)

	app.Get("/login", handlers.LoginGet)
	app.Post("/login", handlers.LoginPost)

	app.Get("/register", handlers.RegisterGet)
	app.Post("/register", handlers.RegisterPost)

	app.Use(middlewares.IsAuthenticated)
	//TODO: Şifre gizlenecek (sor)
	//TODO: Giriş yapan kullanıcı dışında başka kullanıcının bilgileri gözükmeyecek
	//TODO: Aynı kullanıcı kayıt kontrolu yap
	app.Get("/user", handlers.User)
	app.Get("/user/:key", handlers.GetUser)

	app.Get("/admin", handlers.Admin)

	//app.Get("/admin/:key", handlers.GetUser)
	app.Get("/update", handlers.Update)
	app.Get("/update/:key", handlers.GetUser)
	app.Post("/update/:key", handlers.GetUpdate)

	app.Get("/delete/:key", handlers.Delete)
	app.Post("/update/:key", handlers.Delete)

	app.Get("/success", handlers.Logout)

	app.Get("/unsuccess", handlers.Unsuccess)
}
