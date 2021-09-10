package models

import (
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	//	CustomersID uint
	PaymentsID uint `json:"payments_id"`
	//CartsID            uint
	TotalPrice         int `json:"total_price"`
	TransactionsStatus string `json:"transactions_status"`
	TransactionsDate   time.Time `json:"transactions_date"`
	Deliveries         Deliveries
}
