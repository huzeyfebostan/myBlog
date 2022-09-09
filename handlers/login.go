package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/huzeyfebostan/myBlog/database"
	"github.com/huzeyfebostan/myBlog/models"
)

/*var store = session.New(session.Config{
	Expiration:   24 * time.Hour,
	KeyLookup:    "cookie:session_id",
	KeyGenerator: utils.UUID,
})*/

func LoginGet(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func LoginPost(c *fiber.Ctx) error {

	var request models.RequestSignIn
	var user models.User

	err := c.BodyParser(&request)
	if err != nil {
		fmt.Println(err)
	}

	if err = database.DB().Where("email = ? and password = ?", request.Email, request.Password).First(&user).Error; err != nil {
		return c.Redirect("unsuccess")
	}

	/*sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
	}

	sess.Save()

	sess.Set("email", user.Email)
	sess.Set("password", user.Password)*/

	return c.Redirect("/admin")
}

func Unsuccess(c *fiber.Ctx) error {
	return c.Render("unsuccess", fiber.Map{})
}
