package orderservice

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/imnzr/virtual-number-service/helper"
	"github.com/imnzr/virtual-number-service/models"
	orderrepository "github.com/imnzr/virtual-number-service/repository/order_repository"
	"github.com/imnzr/virtual-number-service/utils"
)

type OrderServiceImplement struct {
	OrderRepository orderrepository.OrderRepositoryInterface
	DB              *sql.DB
}

// BuyNumber implements OrderServiceInterface.
func (service *OrderServiceImplement) BuyNumber(ctx context.Context, country string, operator string, product string, userId int) (*models.SMSOrder, error) {
	client := http.Client{}
	url := fmt.Sprintf("%suser/buy/activation/%s/%s/%s", os.Getenv("SIM_API_URL_SERVICE"), country, operator, product)

	req, err := utils.NewRequestSIM("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	fmt.Println("DEBUG BODY:", string(bodyBytes))

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("buy number error: %s", string(bodyBytes))
	}

	var data struct {
		Id     int64  `json:"id"`
		Phone  string `json:"phone"`
		Expiry string `json:"expires"`
	}

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	expiredAt, _ := time.Parse(time.RFC3339, data.Expiry)

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	order := &models.SMSOrder{
		OrderId:   data.Id,
		Phone:     data.Phone,
		Country:   country,
		Operator:  operator,
		Product:   product,
		Status:    "waiting",
		UserId:    userId,
		ExpiredAt: &expiredAt,
	}
	if err := service.OrderRepository.Create(ctx, tx, order); err != nil {
		return nil, err
	}
	return order, nil
}

// CancelOrder implements OrderServiceInterface.
func (service *OrderServiceImplement) CancelOrder(ctx context.Context, orderId int64) error {
	client := &http.Client{}
	url := fmt.Sprintf("%suser/cancel/%d", os.Getenv("SIM_API_URL_SERVICE"), orderId)

	req, err := utils.NewRequestSIM("PUT", url, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(res.Body)
		return fmt.Errorf("cancel order error: %s", string(bodyBytes))
	}
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.OrderRepository.UpdateStatus(ctx, tx, orderId, "cancelled")
}

// CheckOrder implements OrderServiceInterface.
func (service *OrderServiceImplement) CheckOrder(ctx context.Context, orderId int64) (*models.SMSOrder, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%suser/check/%d", os.Getenv("SIM_API_URL_SERVICE"), orderId)

	req, err := utils.NewRequestSIM("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("check order error: %s", string(bodyBytes))
	}

	var data struct {
		Status string `json:"status"`
		SMS    []struct {
			Code string `json:"code"`
		} `json:"sms"`
	}

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, fmt.Errorf("failed to decode check order response: %w", err)
	}
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	if len(data.SMS) > 0 {
		if err := service.OrderRepository.UpdateCode(ctx, tx, orderId, data.SMS[0].Code); err != nil {
			return nil, err
		}
	}

	order, err := service.OrderRepository.FindByOrderId(ctx, tx, orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// FinishOrder implements OrderServiceInterface.
func (service *OrderServiceImplement) FinishOrder(ctx context.Context, orderId int64) error {
	client := &http.Client{}
	url := fmt.Sprintf("%suser/finish/%d", os.Getenv("SIM_API_URL_SERVICE"), orderId)

	req, err := utils.NewRequestSIM("PUT", url, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(res.Body)
		return fmt.Errorf("finish order error: %s", string(bodyBytes))
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return service.OrderRepository.UpdateStatus(ctx, tx, orderId, "finished")
}

func NewOrderService(orderRepository orderrepository.OrderRepositoryInterface, db *sql.DB) OrderServiceInterface {
	return &OrderServiceImplement{
		OrderRepository: orderRepository,
		DB:              db,
	}
}
