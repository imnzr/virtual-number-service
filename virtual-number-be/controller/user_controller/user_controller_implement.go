package usercontroller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/imnzr/virtual-number-service/config"
	"github.com/imnzr/virtual-number-service/models"
	userservice "github.com/imnzr/virtual-number-service/service/user_service"
	"github.com/imnzr/virtual-number-service/web/request"
	"github.com/imnzr/virtual-number-service/web/response"
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
	users, err := uc.UserService.GetAllUsers(controller.Context())
	if err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var userResponse []response.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, response.UserResponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return controller.Status(200).JSON(userResponse)
}

// GetUserByEmail implements UserControllerInterface.
func (uc *UserControllerImplement) GetUserByEmail(controller *fiber.Ctx) error {
	panic("unimplemented")
}

// GetUserById implements UserControllerInterface.
func (uc *UserControllerImplement) GetUserById(controller *fiber.Ctx) error {
	paramsId := controller.Params("id")
	userId, err := strconv.Atoi(paramsId)

	if err != nil || userId <= 0 {
		return controller.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := uc.UserService.GetUserById(controller.Context(), userId)
	if err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := response.UserResponse{
		Id:       result.Id,
		Username: result.Username,
		Email:    result.Email,
	}

	return controller.Status(200).JSON(response)
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

	response := response.UserLoginResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return controller.Status(200).JSON(response)

	// return controller.JSON(fiber.Map{
	// 	"token": token,
	// 	"user": fiber.Map{
	// 		"Id":       user.Id,
	// 		"Username": user.Username,
	// 		"Email":    user.Email,
	// 	},
	// })
}

// LogoutUser implements UserControllerInterface.
func (uc *UserControllerImplement) LogoutUser(controller *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateUserEmail implements UserControllerInterface.
func (uc *UserControllerImplement) UpdateUserEmail(controller *fiber.Ctx) error {
	var user models.User

	idParams := controller.Params("id")
	userId, err := strconv.Atoi(idParams)

	if err != nil || userId <= 0 {
		return controller.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := controller.BodyParser(&user); err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	getUser, err := uc.UserService.GetUserById(controller.Context(), userId)
	if err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	emailUpdated, err := uc.UserService.UpdateUserEmail(controller.Context(), getUser.Id, user.Email)
	if err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := response.UserResponse{
		Id:       emailUpdated.Id,
		Username: emailUpdated.Username,
		Email:    emailUpdated.Email,
	}

	return controller.Status(200).JSON(response)
}

// UpdateUserPassword implements UserControllerInterface.
func (uc *UserControllerImplement) UpdateUserPassword(controller *fiber.Ctx) error {
	var request request.UpdatePasswordRequest

	paramsId := controller.Params("id")
	userId, err := strconv.Atoi(paramsId)

	if err != nil || userId <= 0 {
		return controller.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := controller.BodyParser(&request); err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// GET USER ID
	getUser, err := uc.UserService.GetUserById(controller.Context(), userId)
	if err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// UPDATE PASSWORD
	updatePasswordUser, err := uc.UserService.UpdateUserPassword(controller.Context(), getUser.Id, request)
	if err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return controller.Status(200).JSON(updatePasswordUser)
}

// UpdateUserUsername implements UserControllerInterface.
func (uc *UserControllerImplement) UpdateUserUsername(controller *fiber.Ctx) error {
	var user models.User

	idParams := controller.Params("id")
	userId, err := strconv.Atoi(idParams)

	if err != nil || userId <= 0 {
		return controller.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := controller.BodyParser(&user); err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	getUser, err := uc.UserService.GetUserById(controller.Context(), userId)
	if err != nil || getUser == nil {
		return controller.Status(400).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	updatedUser, err := uc.UserService.UpdateUserUsername(controller.Context(), userId, user.Username)
	if err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return controller.JSON(fiber.Map{
		"username": updatedUser.Username,
	})
}

func NewUserController(userservice userservice.UserServiceInterface) UserControllerInterface {
	return &UserControllerImplement{
		UserService: userservice,
	}
}
