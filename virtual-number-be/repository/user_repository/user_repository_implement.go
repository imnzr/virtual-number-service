package userrepository

import (
	"context"
	"database/sql"
	"log"

	"github.com/imnzr/virtual-number-service/models"
)

type UserRepositoryImplement struct{}

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

// DeleteUser implements UserRepositoryInterface.
func (u *UserRepositoryImplement) DeleteUser(ctx context.Context, tx *sql.Tx, userId int) error {
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
func (u *UserRepositoryImplement) UpdateUserPassword(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error) {
	query := "UPDATE users SET password = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Password, user.Id)
	if err != nil {
		log.Printf("failed to execute query update user password: %v", err)
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

// VerifyEmail implements UserRepositoryInterface.
func (u *UserRepositoryImplement) VerifyEmail(ctx context.Context, tx *sql.Tx, userId int) error {
	panic("unimplemented")
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepositoryImplement{}
}
