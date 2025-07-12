package countryservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/imnzr/virtual-number-service/utils"
)

type CountryServiceImplement struct{}

func NewCountryService() CountryServiceInterface {
	return &CountryServiceImplement{}
}

// GetAvailableCountries implements CountryServiceInterface.
func (c *CountryServiceImplement) GetAvailableCountries() (map[string]interface{}, error) {
	client := http.Client{}
	urlFormat := os.Getenv("SIM_API_URL_SERVICE")

	url := fmt.Sprintf("%s/guest/countries", urlFormat)

	req, err := utils.NewRequestSIM("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, err
}
