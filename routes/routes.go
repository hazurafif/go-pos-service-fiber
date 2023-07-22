package routes

import (
	"go-pos-service-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func routes() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	users := app.Group("users")
	users.Post("/login", controllers.Login)
	users.Post("", controllers.CreateUser)
	users.Delete("/:id", controllers.DeleteUser)
	users.Put("/:id", controllers.UpdateUser)
	users.Get("", controllers.GetAllUsers)
	users.Get("/:id", controllers.GetUserByID)

	product := app.Group("product")
	product.Post("")
	product.Delete("/:id")
	product.Put("/:id")
	product.Get("")
	product.Get("/:id")

	category := app.Group("category")
	category.Post("")
	category.Delete("/:id")
	category.Put("/:id")
	category.Get("")
	category.Get("/:id")

	payment := app.Group("payment")
	payment.Post("")
	payment.Delete("/:id")
	payment.Put("/:id")
	payment.Get("")
	payment.Get("/:id")

	order := app.Group("order")
	order.Post("")
	order.Get("")
	order.Get("/:id")
}
