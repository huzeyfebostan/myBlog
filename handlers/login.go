package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/huzeyfebostan/myBlog/database"
	"github.com/huzeyfebostan/myBlog/middlewares"
	"github.com/huzeyfebostan/myBlog/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

/*var store = session.New(session.Config{
	Expiration:   24 * time.Hour,
	KeyLookup:    "cookie:session_id",
	KeyGenerator: utils.UUID,
})*/

const SecretKey = "secret"

func LoginGet(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func LoginPost(c *fiber.Ctx) error {

	var request models.RequestSignIn
	var user models.User

	/*nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24)*/

	err := c.BodyParser(&request)
	if err != nil {
		fmt.Println(err)
	}

	database.DB().Where("email = ?", request.Email).First(&user)

	if user.Id == 0 {
		return c.Redirect("/unsucces")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return c.Redirect("/unsuccess")
	}

	/*sess, err := store.Get(c)
	if err != nil {
		return err
	}

	sess.Save()

	sess.Set("fname", user.FirstName)
	sess.Set("lname", user.LastName)
	sess.Set("email", user.Email)
	sess.Set("psw", user.Password)*/

	/*clams := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: jwt.NewNumericDate(expireTime),
	})

	token, err := clams.SignedString([]byte(SecretKey))

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)*/

	token, err := middlewares.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return err
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Redirect("/admin")
}

func Unsuccess(c *fiber.Ctx) error {
	return c.Render("unsuccess", fiber.Map{})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	if err := UserControl(c); err != true {
		return c.Redirect("/unsuccess")
	}
	return c.Render("success", fiber.Map{})
}

func UserControl(c *fiber.Ctx) bool {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return false
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	database.DB().Where("id = ?", claims.Issuer).First(&user)

	return true
}

func GetUser(c *fiber.Ctx) error {
	temp := GetUserId(c.Params("key"))

	/*user := models.User{
		Id: temp.Id,
	}*/

	database.DB().Preload("ID ").Find(&temp)

	return c.Render("admin", temp)
}

func GetUserId(id string) models.User {
	var user models.User
	err := database.DB().Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return user
}
