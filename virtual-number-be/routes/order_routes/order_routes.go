package orderroutes

import (
	"github.com/gofiber/fiber/v2"
	ordercontroller "github.com/imnzr/virtual-number-service/controller/order_controller"
)

func OrderRoutes(router *fiber.App, orderController ordercontroller.OrderControllerInterface) {
	router.Post("/api/v1/order/buy", orderController.BuyNumber)
	router.Get("/api/v1/order/check/:orderId", orderController.CheckOrder)
	router.Put("/api/v1/order/finish/:orderId", orderController.FinishOrder)
	router.Put("/api/v1/order/cancel/:orderId", orderController.CancelOrder)
}
