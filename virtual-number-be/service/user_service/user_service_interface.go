package userservice

import (
	"context"

	"github.com/imnzr/virtual-number-service/models"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, request *models.User) error
	DeleteUser(ctx context.Context, user_id int) error

	UpdateUserUsername(ctx context.Context, user_id int, username string) (*models.User, error)
	UpdateUserEmail(ctx context.Context, user_id int) (*models.User, error)
	UpdateUserPassword(ctx context.Context, user_id int) (*models.User, error)

	GetAllUsers(ctx context.Context) ([]*models.User, error)
	GetUserById(ctx context.Context, user_id int) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	LoginUser(ctx context.Context, email string, password string) (*models.User, error)
	LogoutUser(ctx context.Context, user_id int) error

	ForgotPassword(ctx context.Context, email string) (string, error)
}
