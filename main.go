package main

import (
	""
	"models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

	users := app.Group("users")
	users.Post("/login", userscontroller.loginHandler)
	users.Post("")
	users.Delete("/:id")
	users.Put("/:id")
	users.Get("")
	users.Get("/:id")
}
