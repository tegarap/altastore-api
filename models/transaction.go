package models

import (
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	CustomersID        uint
	PaymentsID         uint
	CartsID            uint
	TotalPrice         int
	TransactionsStatus string
	TransactionsDate   time.Time
	Deliveries         Deliveries
}
