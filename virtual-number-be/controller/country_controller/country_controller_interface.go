package countrycontroller

import "github.com/gofiber/fiber/v2"

type CountryControllerInterface interface {
	GetAvailableCountries(controller *fiber.Ctx) error
}
