package models

import (
	"database/sql"
	"time"
)

type SMSOrder struct {
	Id        int            `json:"id"`
	OrderId   int64          `json:"order_id"`
	Phone     string         `json:"phone"`
	Country   string         `json:"country"`
	Operator  string         `json:"operator"`
	Product   string         `json:"product"`
	Code      sql.NullString `json:"code,omitempty"`
	Status    string         `json:"status"`
	ExpiredAt *time.Time     `json:"expired_at"`
	UserId    int            `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
