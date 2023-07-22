package routes

import "github.com/gofiber/fiber/v2"

func routes() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

	users := app.Group("users")
	users.Post("/login")
	users.Post("")
	users.Delete("/:id")
	users.Put("/:id")
	users.Get("")
	users.Get("/:id")

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
