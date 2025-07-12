package countrycontroller

import (
	"github.com/gofiber/fiber/v2"
	countryservice "github.com/imnzr/virtual-number-service/service/country_service"
)

type CountryControllerImplement struct {
	CountryService countryservice.CountryServiceInterface
}

func NewCountryController(countryService countryservice.CountryServiceInterface) CountryControllerInterface {
	return &CountryControllerImplement{
		CountryService: countryService,
	}
}

// GetAvailableCountries implements CountryControllerInterface.
func (c *CountryControllerImplement) GetAvailableCountries(controller *fiber.Ctx) error {
	result, err := c.CountryService.GetAvailableCountries()
	if err != nil {
		return controller.JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	return controller.JSON(result)
}
