package usercontroller

import "github.com/gofiber/fiber/v2"

type UserControllerInterface interface {
	CreateUser(controller *fiber.Ctx) error
	DeleteUser(controller *fiber.Ctx) error

	UpdateUserUsername(controller *fiber.Ctx) error
	UpdateUserEmail(controller *fiber.Ctx) error
	UpdateUserPassword(controller *fiber.Ctx) error

	GetAllUsers(controller *fiber.Ctx) error
	GetUserById(controller *fiber.Ctx) error
	GetUserByEmail(controller *fiber.Ctx) error

	LoginUser(controller *fiber.Ctx) error
	LogoutUser(controller *fiber.Ctx) error

	ForgotPassword(controller *fiber.Ctx) error
	VerifyResetToken(controller *fiber.Ctx) error
}
