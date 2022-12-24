package handlers

import (
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

func Login(c *fiber.Ctx) error {

	var request models.RequestSignIn
	var user models.User
	var log models.UserLog

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	database.DB().Where("email = ?", request.Email).First(&user)

	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Yok yok kullanıcı yok",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Şifren yanlış aloo",
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

	log.UserId = user.Id
	database.DB().Create(&log)
	/*if user.RoleId == 1 {
		return c.Redirect("/admin")
	}*/

	return c.JSON(user)
}

//TODO: Kullanıcı logout yapmadığı zaman serveri tekrar bile başlatsan Authenticated çalışmıyor. Çıkış yaptıktan sonra çalışıyor sadece

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	if err := UserControl(c); err != true {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Giriş yapmadan nasıl çıkış yapıyorsun",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Başarılı",
	})
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

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := middlewares.ParseJwt(cookie)

	var user models.User

	database.DB().Where("id = ?", id).First(&user)

	return c.JSON(user)
}
