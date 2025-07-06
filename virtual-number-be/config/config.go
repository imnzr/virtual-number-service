package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SimApiKeyService string `json:"sim_api_key_service"`
	SimApiUrlService string `json:"sim_api_url_service"`
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	return Config{
		SimApiKeyService: os.Getenv("SIM_API_KEY_SERVICE"),
		SimApiUrlService: os.Getenv("SIM_API_URL_SERVICE"),
	}
}
