package usercontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imnzr/virtual-number-service/config"
	"github.com/imnzr/virtual-number-service/models"
	userservice "github.com/imnzr/virtual-number-service/service/user_service"
)

type UserControllerImplement struct {
	UserService userservice.UserServiceInterface
}

// CreateUser implements UserControllerInterface.
func (uc *UserControllerImplement) CreateUser(controller *fiber.Ctx) error {
	var user models.User

	if err := controller.BodyParser(&user); err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}
	if err := uc.UserService.CreateUser(controller.Context(), &user); err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return controller.Status(200).JSON(fiber.Map{
		"success": "User created successfully",
	})
}

// DeleteUser implements UserControllerInterface.
func (uc *UserControllerImplement) DeleteUser(controller *fiber.Ctx) error {
	panic("unimplemented")
}

// ForgotPassword implements UserControllerInterface.
func (uc *UserControllerImplement) ForgotPassword(controller *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAllUsers implements UserControllerInterface.
func (uc *UserControllerImplement) GetAllUsers(controller *fiber.Ctx) error {
	panic("unimplemented")
}

// GetUserByEmail implements UserControllerInterface.
func (uc *UserControllerImplement) GetUserByEmail(controller *fiber.Ctx) error {
	panic("unimplemented")
}

// GetUserById implements UserControllerInterface.
func (uc *UserControllerImplement) GetUserById(controller *fiber.Ctx) error {
	panic("unimplemented")
}

// LoginUser implements UserControllerInterface.
func (uc *UserControllerImplement) LoginUser(controller *fiber.Ctx) error {
	var input models.User

	if err := controller.BodyParser(&input); err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := uc.UserService.LoginUser(controller.Context(), input.Email, input.Password)
	if err != nil {
		return controller.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := config.GenerateJWT(user.Email)
	if err != nil {
		return controller.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return controller.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"Id":       user.Id,
			"Username": user.Username,
			"Email":    user.Email,
		},
	})
}

// LogoutUser implements UserControllerInterface.
func (uc *UserControllerImplement) LogoutUser(controller *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateUserEmail implements UserControllerInterface.
func (uc *UserControllerImplement) UpdateUserEmail(controller *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateUserPassword implements UserControllerInterface.
func (uc *UserControllerImplement) UpdateUserPassword(controller *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateUserUsername implements UserControllerInterface.
func (uc *UserControllerImplement) UpdateUserUsername(controller *fiber.Ctx) error {
	panic("unimplemented")
}

func NewUserController(userservice userservice.UserServiceInterface) UserControllerInterface {
	return &UserControllerImplement{
		UserService: userservice,
	}
}
