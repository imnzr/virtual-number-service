package orderrepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/imnzr/virtual-number-service/models"
)

type OrderRepositoryImplement struct{}

// Create implements OrderRepositoryInterface.
func (o *OrderRepositoryImplement) Create(ctx context.Context, tx *sql.Tx, order *models.SMSOrder) error {
	query := "INSERT INTO sms_orders(order_id, phone, country, operator, product, status, user_id, created_at, updated_at) VALUES(?,?,?,?,?,?,?, NOW(), NOW())"

	_, err := tx.ExecContext(ctx, query, order.OrderId, order.Phone, order.Country, order.Operator, order.Product, order.Status, order.UserId)

	if err != nil {
		return fmt.Errorf("failed to execute query insert order")
	}

	return err
}

// FindAllByUserId implements OrderRepositoryInterface.
func (o *OrderRepositoryImplement) FindAllByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]models.SMSOrder, error) {
	panic("unimplemented")
}

// FindByOrderId implements OrderRepositoryInterface.
func (o *OrderRepositoryImplement) FindByOrderId(ctx context.Context, tx *sql.Tx, orderId int64) (*models.SMSOrder, error) {
	query := "SELECT id, order_id, phone, country, operator, product, status, user_id, created_at, updated_at WHERE order_id = ?"

	rows, err := tx.QueryContext(ctx, query, orderId)
	if err != nil {
		return nil, fmt.Errorf("failed to query rows find by order id")
	}

	defer rows.Close()

	if rows.Next() {
		order := &models.SMSOrder{}
		err := rows.Scan(&order.Id, &order.OrderId, &order.Phone, &order.Country, &order.Operator, &order.Product, &order.Status, &order.ExpiredAt, &order.UserId, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			log.Printf("failed to scan user row: %v", err)
			return nil, err
		}
		return order, nil
	}
	log.Printf("no order id found")
	return nil, nil
}

// UpdateCode implements OrderRepositoryInterface.
func (o *OrderRepositoryImplement) UpdateCode(ctx context.Context, tx *sql.Tx, orderId int64, code string) error {
	panic("unimplemented")
}

// UpdateStatus implements OrderRepositoryInterface.
func (o *OrderRepositoryImplement) UpdateStatus(ctx context.Context, tx *sql.Tx, orderId int64, status string) error {
	panic("unimplemented")
}

func NewOrderRepository() OrderRepositoryInterface {
	return &OrderRepositoryImplement{}
}
