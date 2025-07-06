package userrepository

import (
	"context"
	"database/sql"

	"github.com/imnzr/virtual-number-service/models"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, tx *sql.Tx, userId int) error

	UpdateUserUsernme(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error)
	UpdateUSerEmail(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error)
	UpdateUserPassword(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error)

	GetAllUsers(ctx context.Context, tx *sql.Tx) ([]*models.User, error)
	GetUserById(ctx context.Context, tx *sql.Tx, userId int) (*models.User, error)
	GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (*models.User, error)

	LoginUser(ctx context.Context, tx *sql.Tx, email string, password string) (*models.User, error)
	LogoutUser(ctx context.Context, tx *sql.Tx, userId int) error

	ForgotPassword(ctx context.Context, tx *sql.Tx, email string) (string, error)
	ChangePassword(ctx context.Context, tx *sql.Tx, userId int, oldPassword string, newPassword string) error
	VerifyEmail(ctx context.Context, tx *sql.Tx, userId int) error
	ResendVerificationEmail(ctx context.Context, tx *sql.Tx, userId int) error
}
