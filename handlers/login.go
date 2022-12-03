package handlers

import (
	"errors"
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

func GetLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func Login(c *fiber.Ctx) error {

	var request models.RequestSignIn
	var user models.User

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	database.DB().Where("email = ?", request.Email).First(&user)

	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "There is no such user",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Wrong password",
		})
	}

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

	/*if user.RoleId == 1 {
		return c.Redirect("/admin")
	}*/

	return c.JSON(user)
}

//TODO: Kullanıcı logout yapmadığı zaman serveri tekrar bile başlatsan Authenticated çalışmıyoru. Çıkış yaptıktan sonra çalışıyor sadece

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	if err := UserControl(c); err != true {
		return errors.New("Giriş yapmadan nasıl çıkış yapıyorsun")
	}
	return c.Redirect("/")
}

//TODO: AllUser fonksiyonu ekle, bütün kullanıcıları listeleyecek fonksiyon

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

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := middlewares.ParseJwt(cookie)

	var user models.User

	database.DB().Where("id = ?", id).First(&user)

	if user.RoleId == 1 {
		return c.Render("admin", user)
	}

	return c.Render("user", user)
}

func GetUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")
	id, _ := middlewares.ParseJwt(cookie)
	Id, _ := strconv.Atoi(id)

	activeUser := GetUserId(id)

	temp := GetUserId(c.Params("key"))

	if activeUser.RoleId == 1 {
		return c.Render("adminUpdate", temp)
	} else {
		if uint(Id) == temp.Id {
			database.DB().Preload("id ").Find(&temp)

			return c.Render("update", temp)
		}

	}
	return errors.New("unauthorized")
}

func GetUserId(id string) models.User {
	var user models.User
	err := database.DB().Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func Update(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var request models.RequestRegister

	id, _ := strconv.Atoi(c.Params("key"))

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if request.Password != request.PasswordConfirm {
		return errors.New("passwords do not match")
	}

	user := models.User{
		Id:        uint(id),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}

	if request.Password != "" {
		user.SetPassword(user.Password)
	}

	database.DB().Model(&user).Where("id = ?", user.Id).Updates(user)

	if user.RoleId == 1 {
		return c.Redirect("/admin")
	}

	return c.Redirect("/user")
}

func Delete(c *fiber.Ctx) error {
	/*key := c.Params("key")

	var deluser models.User
	if err := database.DB().Where("id = ?", key).Delete(&deluser).Error; err != nil {
		return err
	}*/

	id, _ := strconv.Atoi(c.Params("key"))

	user := models.User{
		Id: uint(id),
	}
	database.DB().Where("id = ?", id).First(&user)

	database.DB().Delete(&user)

	return c.Redirect("/")
}
