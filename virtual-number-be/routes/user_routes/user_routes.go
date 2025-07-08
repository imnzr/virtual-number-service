package userroutes

import (
	"github.com/gofiber/fiber/v2"
	usercontroller "github.com/imnzr/virtual-number-service/controller/user_controller"
)

// func UserRoutes(router *httprouter.Router, userController usercontroller.UserControllerInterface) {
// 	router.POST("/api/v1/user/create", userController.CreateUser())
// }

func UserRoutes(router *fiber.App, userController usercontroller.UserControllerInterface) {
	router.Post("/api/v1/user/create", userController.CreateUser)
	router.Post("/api/v1/user/login", userController.LoginUser)

	router.Put("/api/v1/user/update-username/:id", userController.UpdateUserUsername)
}
