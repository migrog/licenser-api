package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/migrog/licenser-api/license/infraestructure/di"
)

func AddLicenseRouter(router fiber.Router) {
	router.Post("/licenses", di.LicenseController.CreateLicense)
	router.Get("/licenses/verify", di.LicenseController.VerifyLicense)
	router.Put("/licenses/enable", di.LicenseController.EnableLicense)
	router.Put("/licenses/disable", di.LicenseController.DisableLicense)
}
