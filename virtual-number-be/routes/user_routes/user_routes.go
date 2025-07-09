package userroutes

import (
	"github.com/gofiber/fiber/v2"
	usercontroller "github.com/imnzr/virtual-number-service/controller/user_controller"
)

func UserRoutes(router *fiber.App, userController usercontroller.UserControllerInterface) {
	router.Post("/api/v1/user/create", userController.CreateUser)
	router.Post("/api/v1/user/login", userController.LoginUser)

	router.Put("/api/v1/user/update-username/:id", userController.UpdateUserUsername)
	router.Put("/api/v1/user/update-email/:id", userController.UpdateUserEmail)
	router.Put("/api/v1/user/update-password/:id", userController.UpdateUserPassword)

	router.Get("/api/v1/user/:id", userController.GetUserById)
	router.Get("/api/v1/users", userController.GetAllUsers)
}
