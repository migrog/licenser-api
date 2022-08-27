package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/migrog/licenser-api/license/domain/dto"
	"github.com/migrog/licenser-api/license/domain/interfaces"
)

type LicenseController struct {
	Service interfaces.ILicenseService
}

func NewLicenseController(s interfaces.ILicenseService) *LicenseController {
	return &LicenseController{Service: s}
}

func (s *LicenseController) CreateLicense(c *fiber.Ctx) error {
	req := new(dto.CreateLicenseRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}
	if err := req.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res := s.Service.CreateLicense(req)
	return c.Status(res.StatusCode).JSON(res.Data)
}

func (s *LicenseController) VerifyLicense(c *fiber.Ctx) error {
	req := new(dto.VerifyLicenseRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}

	if err := req.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res := s.Service.VerifyLicense(req)
	return c.Status(res.StatusCode).JSON(res.Data)
}

func (s *LicenseController) EnableLicense(c *fiber.Ctx) error {
	req := new(dto.EnableLicenseRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}

	if err := req.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res := s.Service.EnableLicense(req)
	return c.Status(res.StatusCode).JSON(res.Data)
}
func (s *LicenseController) DisableLicense(c *fiber.Ctx) error {
	req := new(dto.DisableLicenseRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}

	if err := req.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res := s.Service.DisableLicense(req)
	return c.Status(res.StatusCode).JSON(res.Data)
}
