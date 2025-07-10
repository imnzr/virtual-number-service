package models

import "time"

type ResetPassword struct {
	Id     int
	Email  string
	Token  string
	Expire time.Time
}
