package productcontroller

import (
	"github.com/gofiber/fiber/v2"
	productservice "github.com/imnzr/virtual-number-service/service/product_service"
)

type ProductControllerImplement struct {
	ProductService productservice.ProductServiceInterface
}

func NewProductController(productService productservice.ProductServiceInterface) ProductControllerInterface {
	return &ProductControllerImplement{
		ProductService: productService,
	}
}

// GetProductAvailable implements ProductControllerInterface.
func (p *ProductControllerImplement) GetProductAvailable(controller *fiber.Ctx) error {
	country := controller.Params("country")
	operator := controller.Params("operator")

	if country == "" || operator == "" {
		return controller.Status(400).JSON(fiber.Map{
			"message": "country and operator are required",
		})
	}

	products, err := p.ProductService.GetProductAvailable(country, operator)
	if err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return controller.JSON(fiber.Map{
		"country":  country,
		"operator": operator,
		"products": products,
	})
}
