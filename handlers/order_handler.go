package handlers

import (
	"github.com/JinJaeJee/golang-order-fiber-api/models"
	"github.com/JinJaeJee/golang-order-fiber-api/services"
	"github.com/gofiber/fiber/v2"
)

func HandleOrders(c *fiber.Ctx) error {
	var inputOrders []models.InputOrder
	if err := c.BodyParser(&inputOrders); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	cleanedOrders := services.ProcessOrders(inputOrders)
	return c.JSON(cleanedOrders)
}
