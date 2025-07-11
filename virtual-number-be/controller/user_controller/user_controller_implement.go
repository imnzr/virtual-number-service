package usercontroller

import (
	"log"
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

// VerifyResetToken implements UserControllerInterface.
func (uc *UserControllerImplement) VerifyResetToken(controller *fiber.Ctx) error {
	type Request struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}

	var req Request

	if err := controller.BodyParser(&req); err != nil || req.Email == "" || req.Token == "" {
		return controller.Status(400).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	errReset := uc.UserService.VerifyResetToken(controller.Context(), req.Email, req.Token)
	if errReset != nil {
		return controller.Status(400).JSON(fiber.Map{
			"message": "token verification failed",
		})
	}
	return controller.Status(200).JSON(fiber.Map{
		"message": "token is valid, go to next change password",
	})
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

	var req request.EmailRequest

	// log.Println("Body raw:", string(controller.Body()))

	if err := controller.BodyParser(&req); err != nil {
		log.Println("❌ Gagal parse body:", err)
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error() + err.Error(),
		})
	}

	// log.Printf("📨 Memproses forgot password untuk %s", req.Email)

	if req.Email == "" {
		return controller.Status(400).JSON(fiber.Map{
			"error": "email tidak boleh kosong",
		})
	}

	// log.Printf("📧 Request email: %s", req.Email)

	err := uc.UserService.ForgotPassword(controller.Context(), req.Email)
	if err != nil {
		// log.Println("❌ Error dari service:", err)
		return controller.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// log.Println("✅ ForgotPassword selesai tanpa error")

	return controller.JSON(fiber.Map{
		"message": "Kode verifikasi telah dikirim ke email",
	})
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
	type Request struct {
		Email       string `json:"email"`
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}

	var req Request

	if err := controller.BodyParser(&req); err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"message": "request invalid",
		})
	}

	err := uc.UserService.ResetPassword(controller.Context(), req.Email, req.Token, req.NewPassword)
	if err != nil {
		return controller.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return controller.Status(200).JSON(fiber.Map{
		"message": "Password update successfully",
	})
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
