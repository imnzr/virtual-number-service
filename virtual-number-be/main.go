package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/imnzr/virtual-number-service/config"
	countrycontroller "github.com/imnzr/virtual-number-service/controller/country_controller"
	usercontroller "github.com/imnzr/virtual-number-service/controller/user_controller"
	"github.com/imnzr/virtual-number-service/database"
	userrepository "github.com/imnzr/virtual-number-service/repository/user_repository"
	countryroutes "github.com/imnzr/virtual-number-service/routes/country_routes"
	userroutes "github.com/imnzr/virtual-number-service/routes/user_routes"
	countryservice "github.com/imnzr/virtual-number-service/service/country_service"
	userservice "github.com/imnzr/virtual-number-service/service/user_service"
)

func main() {
	// Load environment variables and database connection
	config.LoadEnv()

	// Redis
	config.InitRedis()

	db, err := database.DatabaseConnection()
	if err != nil {
		log.Fatal("error connecting to the database:", err)
	}

	cfg := config.LoadConfig()
	if cfg.SimApiKeyService == "" || cfg.SimApiUrlService == "" {
		log.Fatal("SIM_API_KEY_SERVICE or SIM_API_URL_SERVICE is not set in the environment variables")
	}

	defer db.Close()

	userrepository := userrepository.NewUserRepository()
	userservice := userservice.NewUserService(userrepository, db)
	usercontroller := usercontroller.NewUserController(userservice)

	countryService := countryservice.NewCountryService()
	countryController := countrycontroller.NewCountryController(countryService)

	// Initialize the router and define routes
	router := fiber.New()

	// Define user routes
	userroutes.UserRoutes(router, usercontroller)
	countryroutes.CountryRoutes(router, countryController)

	// Server listening configuration
	router.Listen("localhost:8080")
}
