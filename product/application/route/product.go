package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/migrog/licenser-api/product/infraestructure/di"
)

func AddLicenseRouter(router fiber.Router) {
	router.Post("/products", di.ProductController.CreateProduct)
	router.Get("/products/:id", di.ProductController.ProductById)
}
