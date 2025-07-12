package productservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/imnzr/virtual-number-service/utils"
)

type ProductServiceImplement struct{}

type ProductInfo struct {
	Category string  `json:"category"`
	Qty      int     `json:"qty"`
	Price    float64 `json:"price"`
}

func NewProductService() ProductServiceInterface {
	return &ProductServiceImplement{}
}

// GetProductAvailable implements ProductServiceInterface.
func (p *ProductServiceImplement) GetProductAvailable(country, operator string) (map[string]ProductInfo, error) {
	url := fmt.Sprintf("%sguest/products/%s/%s", os.Getenv("SIM_API_URL_SERVICE"), country, operator)

	req, err := utils.NewRequestSIM("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)
	bodyText := string(bodyBytes)

	if !strings.HasPrefix(bodyText, "{") {
		return nil, fmt.Errorf("sim service error: %s", bodyText)
	}

	var result map[string]ProductInfo

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to decode product JSON: %w", err)
	}

	return result, nil
}
