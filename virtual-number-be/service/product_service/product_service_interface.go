package productservice

type ProductServiceInterface interface {
	GetProductAvailable(country, operator string) (map[string]ProductInfo, error)
}
