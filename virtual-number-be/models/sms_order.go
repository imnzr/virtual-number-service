package models

import "time"

type SMSOrder struct {
	Id        int
	OrderId   int64
	Phone     string
	Country   string
	Operator  string
	Product   string
	Code      string
	Status    string
	ExpiredAt *time.Time
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
