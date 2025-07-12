package countryroutes

import (
	"github.com/gofiber/fiber/v2"
	countrycontroller "github.com/imnzr/virtual-number-service/controller/country_controller"
)

func CountryRoutes(router *fiber.App, countryController countrycontroller.CountryControllerInterface) {
	router.Get("/api/v1/country", countryController.GetAvailableCountries)
}
