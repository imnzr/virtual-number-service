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
	query := "SELECT id, order_id, phone, country, operator, product, status, user_id, created_at, updated_at WHERE id = ? ORDER BY created_at DESC"
	rows, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []models.SMSOrder

	for rows.Next() {
		var order models.SMSOrder
		err := rows.Scan(&order.Id, &order.OrderId, &order.Phone, &order.Country, &order.Operator, &order.Product, &order.Code, &order.Status, &order.ExpiredAt, &order.UserId, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	return orders, nil
}

// FindByOrderId implements OrderRepositoryInterface.
func (o *OrderRepositoryImplement) FindByOrderId(ctx context.Context, tx *sql.Tx, orderId int64) (*models.SMSOrder, error) {
	query := `
		SELECT id, order_id, phone, country, operator, product, code, status, expired_at, user_id, created_at, updated_at
		FROM sms_orders
		WHERE order_id = ?
	`

	row := tx.QueryRowContext(ctx, query, orderId)

	order := &models.SMSOrder{}
	err := row.Scan(
		&order.Id,
		&order.OrderId,
		&order.Phone,
		&order.Country,
		&order.Operator,
		&order.Product,
		&order.Code,
		&order.Status,
		&order.ExpiredAt,
		&order.UserId,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Order ID %d tidak ditemukan", orderId)
			return nil, fmt.Errorf("order id %d tidak ditemukan", orderId)
		}
		log.Printf("Gagal scan baris: %v", err)
		return nil, fmt.Errorf("failed to scan order row: %w", err)
	}

	return order, nil
}

// UpdateCode implements OrderRepositoryInterface.
func (r *OrderRepositoryImplement) UpdateCode(ctx context.Context, tx *sql.Tx, orderId int64, code string) error {
	query := `UPDATE sms_orders SET code = ?, updated_at = NOW() WHERE order_id = ?`
	_, err := tx.ExecContext(ctx, query, code, orderId)
	return err
}

// UpdateStatus implements OrderRepositoryInterface.
func (r *OrderRepositoryImplement) UpdateStatus(ctx context.Context, tx *sql.Tx, orderId int64, status string) error {
	query := `UPDATE sms_orders SET status = ?, updated_at = NOW() WHERE order_id = ?`
	_, err := tx.ExecContext(ctx, query, status, orderId)
	return err
}

func NewOrderRepository() OrderRepositoryInterface {
	return &OrderRepositoryImplement{}
}
