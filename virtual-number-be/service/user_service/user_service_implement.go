package userservice

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/imnzr/virtual-number-service/helper"
	"github.com/imnzr/virtual-number-service/models"
	userrepository "github.com/imnzr/virtual-number-service/repository/user_repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImplement struct {
	UserRepository userrepository.UserRepositoryInterface
	DB             *sql.DB
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
func (service *UserServiceImplement) ForgotPassword(ctx context.Context, email string) (string, error) {
	panic("unimplemented")
}

// GetAllUsers implements UserServiceInterface.
func (service *UserServiceImplement) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements UserServiceInterface.
func (service *UserServiceImplement) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	panic("unimplemented")
}

// GetUserById implements UserServiceInterface.
func (service *UserServiceImplement) GetUserById(ctx context.Context, user_id int) (*models.User, error) {
	panic("unimplemented")
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
func (service *UserServiceImplement) UpdateUserEmail(ctx context.Context, user_id int) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUserPassword implements UserServiceInterface.
func (service *UserServiceImplement) UpdateUserPassword(ctx context.Context, user_id int) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUserUsername implements UserServiceInterface.
func (service *UserServiceImplement) UpdateUserUsername(ctx context.Context, user_id int) (*models.User, error) {
	panic("unimplemented")
}

func NewUserService(userRepository userrepository.UserRepositoryInterface, db *sql.DB) UserServiceInterface {
	return &UserServiceImplement{
		UserRepository: userRepository,
		DB:             db,
	}
}
