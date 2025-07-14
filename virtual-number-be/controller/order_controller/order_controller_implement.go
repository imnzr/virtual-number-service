package ordercontroller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	orderservice "github.com/imnzr/virtual-number-service/service/order_service"
)

type OrderController struct {
	OrderService orderservice.OrderServiceInterface
}

// CheckOrder implements OrderControllerInterface.
func (o *OrderController) CheckOrder(c *fiber.Ctx) error {
	orderIdParams := c.Params("orderId")
	orderId, err := strconv.ParseInt(orderIdParams, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	order, err := o.OrderService.CheckOrder(c.Context(), orderId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(order)
}

// BuyNumber implements OrderControllerInterface.
func (o *OrderController) BuyNumber(c *fiber.Ctx) error {
	type BuyRequest struct {
		Country  string `json:"country"`
		Operator string `json:"operator"`
		Product  string `json:"product"`
		UserID   int    `json:"id"`
	}

	var req BuyRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}
	order, err := o.OrderService.BuyNumber(c.Context(), req.Country, req.Operator, req.Product, req.UserID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(order)
}

// CancelOrder implements OrderControllerInterface.
func (o *OrderController) CancelOrder(c *fiber.Ctx) error {
	orderIdParams := c.Params("orderId")
	orderId, err := strconv.ParseInt(orderIdParams, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid order id",
		})
	}

	if err := o.OrderService.CancelOrder(c.Context(), orderId); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "order cancelled",
	})
}

// FinishOrder implements OrderControllerInterface.
func (o *OrderController) FinishOrder(c *fiber.Ctx) error {
	orderIdParams := c.Params("orderId")
	orderId, err := strconv.ParseInt(orderIdParams, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid order id",
		})
	}

	order, err := o.OrderService.CheckOrder(c.Context(), orderId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messaga": err.Error(),
		})
	}

	return c.JSON(order)
}

func NewOrderController(orderService orderservice.OrderServiceInterface) OrderControllerInterface {
	return &OrderController{
		OrderService: orderService,
	}
}
