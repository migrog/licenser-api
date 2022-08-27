package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/migrog/licenser-api/product/domain/dto"
	"github.com/migrog/licenser-api/product/domain/interfaces"
)

type ProductController struct {
	Service interfaces.IProductService
}

func NewProductController(s interfaces.IProductService) *ProductController {
	return &ProductController{Service: s}
}

func (s *ProductController) CreateProduct(c *fiber.Ctx) error {
	req := new(dto.CreateProductRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}
	if err := req.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res := s.Service.CreateProduct(req)
	return c.Status(res.StatusCode).JSON(res.Data)
}

func (s *ProductController) ProductById(c *fiber.Ctx) error {
	res := s.Service.ProductById(c.Params("id"))
	return c.Status(res.StatusCode).JSON(res.Data)
}
