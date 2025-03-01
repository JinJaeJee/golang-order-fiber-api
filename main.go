package main

import (
	"fmt"

	"github.com/JinJaeJee/golang-order-fiber-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Println("Hello, World!")
	app.Post("/orders", handlers.HandleOrders)

	app.Listen(":3333")

}
