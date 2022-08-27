package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/migrog/licenser-api/security/domain"
)

type UserHandler struct {
	UService domain.UserService
}

func NewUserHandler(api fiber.Router, us domain.UserService) {
	handler := &UserHandler{
		UService: us,
	}
	api.Get("/users/:email", handler.UserByEmail)
}

func (s *UserHandler) UserByEmail(c *fiber.Ctx) error {
	user, _ := s.UService.GetByEmail(c.Params("email"))

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(user)
}
