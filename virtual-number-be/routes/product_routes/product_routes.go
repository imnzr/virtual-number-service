package productroutes

import (
	"github.com/gofiber/fiber/v2"
	productcontroller "github.com/imnzr/virtual-number-service/controller/product_controller"
)

func ProductRoutes(routes *fiber.App, productController productcontroller.ProductControllerInterface) {
	routes.Get("/api/v1/products/:country/:operator", productController.GetProductAvailable)
}
