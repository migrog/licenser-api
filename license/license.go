package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/migrog/licenser-api/license/application/route"
	"github.com/migrog/licenser-api/license/infraestructure/database"
	"github.com/migrog/licenser-api/license/infraestructure/di"
	"github.com/migrog/licenser-api/license/infraestructure/enviroment"
)

const defaultPort = "3002"
const defaultEnv = "development"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = defaultEnv
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	enviroment.NewSettings(env)
	mongoConnection := database.NewConnection()
	defer mongoConnection.Disconnect()

	di.SetupDependencyInjection(mongoConnection)
	route.SetupRoute(app)

	fmt.Printf("🚀 License API is running on %v", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
