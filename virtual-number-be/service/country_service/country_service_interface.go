package countryservice

type CountryServiceInterface interface {
	GetAvailableCountries() (map[string]interface{}, error)
}
