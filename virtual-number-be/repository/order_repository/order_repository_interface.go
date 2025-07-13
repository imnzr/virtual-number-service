package orderrepository

import (
	"context"
	"database/sql"

	"github.com/imnzr/virtual-number-service/models"
)

type OrderRepositoryInterface interface {
	Create(ctx context.Context, tx *sql.Tx, order *models.SMSOrder) error
	UpdateStatus(ctx context.Context, tx *sql.Tx, orderId int64, status string) error
	UpdateCode(ctx context.Context, tx *sql.Tx, orderId int64, code string) error
	FindByOrderId(ctx context.Context, tx *sql.Tx, orderId int64) (*models.SMSOrder, error)
	FindAllByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]models.SMSOrder, error)
}
