package userservice

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/imnzr/virtual-number-service/helper"
	"github.com/imnzr/virtual-number-service/models"
	userrepository "github.com/imnzr/virtual-number-service/repository/user_repository"
	"github.com/imnzr/virtual-number-service/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImplement struct {
	UserRepository userrepository.UserRepositoryInterface
	DB             *sql.DB
}

// ResetPassword implements UserServiceInterface.
func (service *UserServiceImplement) ResetPassword(ctx context.Context, email string, token string, newPassword string) error {
	tx, err := service.DB.Begin()
	helper.ErrorTransaction(err)
	defer helper.CommitOrRollback(tx)

	resetToken, err := service.UserRepository.FindResetToken(ctx, tx, email, token)
	if err != nil {
		return errors.New("token invalid")
	}
	if time.Now().After(resetToken.Expire) {
		return errors.New("token expired")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed hashed password")
	}

	if err := service.UserRepository.UpdateUserPassword(ctx, tx, email, string(hashedPassword)); err != nil {
		return err
	}

	return service.UserRepository.DeleteResetToken(ctx, tx, email, token)
}

// VerifyResetToken implements UserServiceInterface.
func (service *UserServiceImplement) VerifyResetToken(ctx context.Context, email string, token string) (bool, error) {
	tx, err := service.DB.Begin()
	helper.ErrorTransaction(err)
	defer helper.CommitOrRollback(tx)

	resetToken, err := service.UserRepository.FindResetToken(ctx, tx, email, token)
	if err != nil {
		return false, err
	}
	if resetToken == nil || time.Now().After(resetToken.Expire) {
		return false, err
	}

	return true, nil
}

// CreateUser implements UserServiceInterface.
func (service *UserServiceImplement) CreateUser(ctx context.Context, request *models.User) error {
	tx, err := service.DB.Begin()
	helper.ErrorTransaction(err)
	defer helper.CommitOrRollback(tx)

	// Find email
	existUser, _ := service.UserRepository.GetUserByEmail(ctx, tx, request.Email)
	if existUser != nil {
		log.Printf("email already exist")
		return errors.New("email already exist")
	}

	// hash the password
	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		log.Printf("error hash password")
		return err
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	_, err = service.UserRepository.CreateUser(ctx, tx, &user)
	if err != nil {
		log.Printf("failed to save user: %v", err)
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

// DeleteUser implements UserServiceInterface.
func (service *UserServiceImplement) DeleteUser(ctx context.Context, user_id int) error {
	panic("unimplemented")
}

// ForgotPassword implements UserServiceInterface.
func (service *UserServiceImplement) ForgotPassword(ctx context.Context, email string) error {
	tx, err := service.DB.Begin()
	helper.ErrorTransaction(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.GetUserByEmail(ctx, tx, email)
	if err != nil || user == nil {
		log.Printf("email not found")
		return errors.New("email not found")
	}

	log.Printf("email ditemukan", email)

	token := helper.GenerateToken(6)
	expires := time.Now().Add(5 * time.Minute)

	err = service.UserRepository.SavePasswordReset(ctx, tx, email, token, expires)
	if err != nil {
		log.Printf("error save password reset service")
		return errors.New("error save password reset ")
	}
	body := fmt.Sprintf("kode verifikasi anda adalah: %s. berlaku selama 5 menit", token)

	err = utils.SendEmail(email, "Reset Password", body)
	if err != nil {
		log.Printf("error kirim email: %v", err)
		return errors.New("gagal kirim email")
	}

	log.Printf("email berasil dikirim")

	return nil

}

// GetAllUsers implements UserServiceInterface.
func (service *UserServiceImplement) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	tx, err := service.DB.Begin()
	helper.ErrorTransaction(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.UserRepository.GetAllUsers(ctx, tx)
	if err != nil {
		log.Printf("error to get all users")
		return nil, fmt.Errorf("error to get all users")
	}
	return result, nil
}

// GetUserByEmail implements UserServiceInterface.
func (service *UserServiceImplement) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	tx, err := service.DB.Begin()
	helper.ErrorTransaction(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.UserRepository.GetUserByEmail(ctx, tx, email)
	if err != nil {
		log.Printf("error get user by email")
		return nil, fmt.Errorf("error get user by email")
	}

	return result, err
}

// GetUserById implements UserServiceInterface.
func (service *UserServiceImplement) GetUserById(ctx context.Context, user_id int) (*models.User, error) {
	tx, err := service.DB.Begin()
	helper.ErrorTransaction(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.UserRepository.GetUserById(ctx, tx, user_id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return result, nil
}

// LoginUser implements UserServiceInterface.
func (service *UserServiceImplement) LoginUser(ctx context.Context, email string, password string) (*models.User, error) {
	tx, err := service.DB.Begin()
	helper.ErrorTransaction(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.GetUserByEmail(ctx, tx, email)
	if err != nil {
		log.Printf("error get user by email: %v", err)
		return nil, fmt.Errorf("invalid email or password")
	}
	if user == nil {
		log.Printf("user not found with email: %s", email)
		return nil, fmt.Errorf("invalid email or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("error compare hashing password")
		return nil, fmt.Errorf("invalid email or password")
	}
	return user, nil
}

// LogoutUser implements UserServiceInterface.
func (service *UserServiceImplement) LogoutUser(ctx context.Context, user_id int) error {
	panic("unimplemented")
}

// UpdateUserEmail implements UserServiceInterface.
func (service *UserServiceImplement) UpdateUserEmail(ctx context.Context, user_id int, email string) (*models.User, error) {
	tx, err := service.DB.Begin()
	helper.ErrorTransaction(err)
	defer helper.CommitOrRollback(tx)

	// find by id
	user, err := service.UserRepository.GetUserById(ctx, tx, user_id)
	if err != nil {
		log.Printf("user not found")
		return nil, fmt.Errorf("user not found")
	}

	user.Email = email

	_, err = service.UserRepository.UpdateUserEmail(ctx, tx, user)
	if err != nil {
		log.Printf("error update user email")
		return nil, fmt.Errorf("error update user email")
	}

	return user, nil
}

// // UpdateUserPassword implements UserServiceInterface.
// func (service *UserServiceImplement) UpdateUserPassword(ctx context.Context, user_id int, request request.UpdatePasswordRequest) (*models.User, error) {
// 	tx, err := service.DB.Begin()
// 	helper.ErrorTransaction(err)
// 	defer helper.CommitOrRollback(tx)

// 	// FIND BY ID
// 	user, err := service.UserRepository.GetUserById(ctx, tx, user_id)
// 	if err != nil {
// 		log.Printf("error, user not found")
// 		return nil, fmt.Errorf("error, user not found")
// 	}

// 	// VALIDATE CURRENT PASSWORD
// 	if request.CurrentPassword != "" {
// 		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.CurrentPassword)); err != nil {
// 			log.Printf("invalid current password")
// 			return nil, fmt.Errorf("invalid current password")
// 		}
// 	}

// 	// VALIDATE NEW PASSWORD AND CONFIRMATION
// 	if request.NewPassword != request.ConfirmPassword {
// 		log.Printf("New password and confirmation do not match")
// 		return nil, fmt.Errorf("new password and confirmation do not match")
// 	}

// 	// HASHED NEW PASSWORD
// 	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Printf("error hashed password")
// 		return nil, fmt.Errorf("error hashed password")
// 	}

// 	user.Password = string(hashedNewPassword)

// 	_, err = service.UserRepository.UpdateUserPassword()
// 	if err != nil {
// 		log.Printf("invalid update user password")
// 		return nil, fmt.Errorf("invalid update user password")
// 	}

// 	return user, nil
// }

// UpdateUserUsername implements UserServiceInterface.
func (service *UserServiceImplement) UpdateUserUsername(ctx context.Context, user_id int, username string) (*models.User, error) {
	tx, err := service.DB.Begin()
	helper.ErrorTransaction(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.GetUserById(ctx, tx, user_id)
	if err != nil {
		log.Printf("user not found")
		return nil, fmt.Errorf("error, user not found")
	}

	user.Username = username

	_, err = service.UserRepository.UpdateUserUsername(ctx, tx, user)
	if err != nil {
		log.Printf("error updating user username with id %v", err)
		return nil, fmt.Errorf("error updating username")
	}

	return user, nil
}

func NewUserService(userRepository userrepository.UserRepositoryInterface, db *sql.DB) UserServiceInterface {
	return &UserServiceImplement{
		UserRepository: userRepository,
		DB:             db,
	}
}
