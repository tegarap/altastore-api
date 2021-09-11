package models

import (
	"gorm.io/gorm"
	"time"
)

type Transactions struct {
	gorm.Model
	CustomersID        uint      `json:"customer_id"`
	PaymentsID         uint      `json:"payments_id"`
	CartsID            uint      `json:"carts_id"`
	TotalPrice         int       `json:"total_price"`
	TransactionsStatus string    `json:"transactions_status"`
	TransactionsDate   time.Time `json:"transactions_date"`
	DeliveryStatus     string    `json:"delivery_status" sql:"type:ENUM('pending', 'delivered')"`
}

type TransactionResponse struct {
	ID                 uint
	TotalPrice         int
	TransactionsStatus string
	TransactionsDate   time.Time
	DeliveryStatus     string
}