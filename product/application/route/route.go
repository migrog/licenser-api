package route

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/api")
	AddLicenseRouter(api)

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Operational",
			"time":    time.Now(),
		})
	})
}
