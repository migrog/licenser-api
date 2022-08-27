package http

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/migrog/licenser-api/security/domain"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UService domain.UserService
}

var SecretKey = viper.GetString(`jwt.secret`)

func NewAuthHandler(api fiber.Router, us domain.UserService) {
	handler := &AuthHandler{
		UService: us,
	}

	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)
	api.Get("/user", handler.User)
	api.Get("/logout", handler.Logout)
}

func (a *AuthHandler) Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := domain.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	a.UService.Create(&user)

	return c.JSON(user)
}
func (a *AuthHandler) Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := a.UService.GetByEmail(data["email"])

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func (a *AuthHandler) User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unautheticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	id, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}

	user, err := a.UService.GetById(int64(id))

	return c.JSON(user)
}

func (a *AuthHandler) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
