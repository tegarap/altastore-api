package models

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	PaymentName  string `json:"payment_name"`
	Transactions []Transactions
}
