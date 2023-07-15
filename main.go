package main

import (
	"golang/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	users := app.Group("/users")
	product := app.Group("/product")
	categories := app.Group("/categories")
	payment := app.Group("/payment")
	order := app.Group("/order")

	users.Get("/", userscontroller.Index)
	users.Get(":id", userscontroller.Show)
	users.Post("/", userscontroller.Create)
	users.Put("/:id", userscontroller.Update)
	users.Delete(":id", userscontroller.Delete)
}
