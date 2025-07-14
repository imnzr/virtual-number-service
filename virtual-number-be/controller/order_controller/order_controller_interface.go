package ordercontroller

import "github.com/gofiber/fiber/v2"

type OrderControllerInterface interface {
	BuyNumber(c *fiber.Ctx) error
	FinishOrder(c *fiber.Ctx) error
	CancelOrder(c *fiber.Ctx) error
	CheckOrder(c *fiber.Ctx) error
}
