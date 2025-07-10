package userrepository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/imnzr/virtual-number-service/models"
)

type UserRepositoryImplement struct{}

// DeleteResetToken implements UserRepositoryInterface.
func (u *UserRepositoryImplement) DeleteResetToken(ctx context.Context, tx *sql.Tx, email string, token string) error {
	query := "DELETE FROM password_reset WHERE email = ? AND token = ?"
	_, err := tx.ExecContext(ctx, query, email, token)
	return err
}

// FindResetToken implements UserRepositoryInterface.
func (u *UserRepositoryImplement) FindResetToken(ctx context.Context, tx *sql.Tx, email string, token string) (*models.ResetPassword, error) {
	query := "SELECT email, token, expires_at FROM reset_password WHERE email = ? AND token = ?"

	var reset models.ResetPassword

	row := tx.QueryRowContext(ctx, query)
	err := row.Scan(&reset.Email, &reset.Token, &reset.Expire)
	if err != nil {
		return nil, err
	}

	return &reset, nil
}

// SavePasswordReset implements UserRepositoryInterface.
func (u *UserRepositoryImplement) SavePasswordReset(ctx context.Context, tx *sql.Tx, email, token string, expires time.Time) error {
	query := "INSERT INTO password_reset(email, token, expires_at) VALUES(?,?,?)"
	_, err := tx.ExecContext(ctx, query, email, token, expires)
	if err != nil {
		log.Printf("db error insert password reset: %v", err)
	}
	return err
}

// CreateUser implements UserRepositoryInterface.
func (u *UserRepositoryImplement) CreateUser(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error) {
	query := "INSERT INTO users(username, email, password) VALUES(?,?,?)"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		log.Println("failed to execute query create user: %w", err)
		return nil, err
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println("failed to get last insert id: %w", err)
		return nil, err
	}

	user.Id = int(LastInsertId)
	return user, nil
}

// ChangePassword implements UserRepositoryInterface.
func (u *UserRepositoryImplement) ChangePassword(ctx context.Context, tx *sql.Tx, userId int, oldPassword string, newPassword string) error {
	panic("unimplemented")
}

// ForgotPassword implements UserRepositoryInterface.
func (u *UserRepositoryImplement) ForgotPassword(ctx context.Context, tx *sql.Tx, email string) (string, error) {
	panic("unimplemented")
}

// GetAllUsers implements UserRepositoryInterface.
func (u *UserRepositoryImplement) GetAllUsers(ctx context.Context, tx *sql.Tx) ([]*models.User, error) {
	query := "SELECT id, username, email, password FROM `users`"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		log.Printf("failed to execute query get all users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			log.Printf("failed to scan user row: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByEmail implements UserRepositoryInterface.
func (u *UserRepositoryImplement) GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM `users` WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		log.Println("failed to execute query get user by email: %w", err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			log.Printf("failed to scan user row: %v", err)
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

// GetUserById implements UserRepositoryInterface.
func (u *UserRepositoryImplement) GetUserById(ctx context.Context, tx *sql.Tx, userId int) (*models.User, error) {
	query := "SELECT id, username, email, password FROM `users` WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		log.Printf("failed to execute query get user by id: %v", err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			log.Printf("failed to scan user row: %v", err)
			return nil, err
		}
		return user, nil
	}
	log.Printf("no user found with id: %d", userId)
	return nil, nil
}

// LoginUser implements UserRepositoryInterface.
func (u *UserRepositoryImplement) LoginUser(ctx context.Context, tx *sql.Tx, email string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM `users` WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		log.Printf("failed to execute query login user: %v", err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			log.Printf("failed to scan user row: %v", err)
			return nil, err
		}
		return user, nil
	}
	log.Printf("user not found with email: %s", email)
	return nil, nil
}

// LogoutUser implements UserRepositoryInterface.
func (u *UserRepositoryImplement) LogoutUser(ctx context.Context, tx *sql.Tx, userId int) error {
	panic("unimplemented")
}

// ResendVerificationEmail implements UserRepositoryInterface.
func (u *UserRepositoryImplement) ResendVerificationEmail(ctx context.Context, tx *sql.Tx, userId int) error {
	panic("unimplemented")
}

// UpdateUSerEmail implements UserRepositoryInterface.
func (u *UserRepositoryImplement) UpdateUserEmail(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error) {
	query := "UPDATE users SET email = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Email, user.Id)
	if err != nil {
		log.Printf("failed to execute query update user email: %v", err)
		return nil, err
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to get rows affected: %v", err)
		return nil, err
	}
	if RowsAffected == 0 {
		log.Printf("no user found with id: %d", user.Id)
	}

	return user, nil
}

// UpdateUserPassword implements UserRepositoryInterface.
func (u *UserRepositoryImplement) UpdateUserPassword(ctx context.Context, tx *sql.Tx, email, hashedPassword string) error {
	query := "UPDATE users SET password = ? WHERE email = ?"
	_, err := tx.ExecContext(ctx, query, hashedPassword, email)
	return err
}

// UpdateUserUsernme implements UserRepositoryInterface.
func (u *UserRepositoryImplement) UpdateUserUsername(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error) {
	query := "UPDATE users SET username = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Id)
	if err != nil {
		log.Printf("failed to execute query update user username: %v", err)
		return nil, err
	}
	if result == nil {
		log.Printf("no user found with id: %d", user.Id)
	}

	return user, nil
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepositoryImplement{}
}
