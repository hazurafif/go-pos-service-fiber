package main

import (
	"go-pos-service-fiber/models"
	"go-pos-service-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	models.ConnectDatabase()
}

func main() {
	app := fiber.New()

	routes.Routes(app)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
