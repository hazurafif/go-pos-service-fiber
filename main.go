package main

import (
	"go-pos-service-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func init() {
	models.ConnectDatabase()
}

func main() {
	app := fiber.New()
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
