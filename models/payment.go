package models

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	PaymentName  string
	Transactions []Transactions
}
