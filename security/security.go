package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_userHandler "github.com/migrog/licenser-api/security/application/http"
	database "github.com/migrog/licenser-api/security/repository"
	_userRepository "github.com/migrog/licenser-api/security/repository/mysql"
	_userService "github.com/migrog/licenser-api/security/service"

	"github.com/spf13/viper"

	jwtware "github.com/gofiber/jwt/v3"
)

const defaultPort = "3000"
const defaultEnv = "development"

func init() {
	var err error

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = defaultEnv
	}

	viper.SetConfigName(fmt.Sprintf("setting.%s", env))
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.Connect(viper.GetString(`database.driver`))

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	ur := _userRepository.NewUserRepository()
	us := _userService.NewUserService(ur)

	api := app.Group("/api")
	_userHandler.NewAuthHandler(api, us)
	_userHandler.NewUserHandler(api, us)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    viper.GetString(`jwt.secret`),
	}))

	fmt.Printf("🚀 Authentication API is running on %v", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
