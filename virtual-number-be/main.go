package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/imnzr/virtual-number-service/config"
	usercontroller "github.com/imnzr/virtual-number-service/controller/user_controller"
	"github.com/imnzr/virtual-number-service/database"
	userrepository "github.com/imnzr/virtual-number-service/repository/user_repository"
	userroutes "github.com/imnzr/virtual-number-service/routes/user_routes"
	userservice "github.com/imnzr/virtual-number-service/service/user_service"
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

	userrepository := userrepository.NewUserRepository()
	userservice := userservice.NewUserService(userrepository, db)
	usercontroller := usercontroller.NewUserController(userservice)

	// Initialize the router and define routes
	router := fiber.New()

	// Define user routes
	userroutes.UserRoutes(router, usercontroller)

	// Server listening configuration
	router.Listen("localhost:8080")
}
