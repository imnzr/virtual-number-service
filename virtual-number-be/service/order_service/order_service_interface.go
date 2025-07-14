package orderservice

import (
	"context"

	"github.com/imnzr/virtual-number-service/models"
)

type OrderServiceInterface interface {
	BuyNumber(ctx context.Context, country, operator, product string, userId int) (*models.SMSOrder, error)
	CheckOrder(ctx context.Context, orderId int64) (*models.SMSOrder, error)
	FinishOrder(ctx context.Context, orderId int64) error
	CancelOrder(ctx context.Context, orderId int64) error
}
