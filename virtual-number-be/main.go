package main

import (
	"log"
	"net/http"

	"github.com/imnzr/virtual-number-service/config"
	"github.com/imnzr/virtual-number-service/database"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Load environment variables and database connection
	config.LoadEnv()
	db, err := database.DatabaseConnection()
	if err != nil {
		log.Fatal("error connecting to the database:", err)
	}

	cfg := config.LoadConfig()
	if cfg.SimApiKeyService == "" || cfg.SimApiUrlService == "" {
		log.Fatal("SIM_API_KEY_SERVICE or SIM_API_URL_SERVICE is not set in the environment variables")
	}

	defer db.Close()

	// Initialize the router and define routes
	router := httprouter.New()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Println("failed to start server:", err)
		return
	}

}
