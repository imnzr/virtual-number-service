package productcontroller

import "github.com/gofiber/fiber/v2"

type ProductControllerInterface interface {
	GetProductAvailable(controller *fiber.Ctx) error
}
