package models

type SMSOrder struct {
	Id     int     `json:"id"`
	Phone  string  `json:"phone"`
	Status string  `json:"status"`
	Price  float64 `json:"price"`
}

type SMSMessage struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type SMSCheckResult struct {
	Status string       `json:"status"`
	SMS    []SMSMessage `json:"sms"`
}
