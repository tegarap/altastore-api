package models

import (
	"time"
)

type Transactions struct {
	ID                 uint       `gorm:"primarykey"`
	CustomersID        uint       `json:"customer_id"`
	PaymentsID         uint       `json:"payments_id"`
	CartsID            uint       `json:"carts_id"`
	TotalPrice         int        `json:"total_price"`
	TransactionsStatus string     `json:"transactions_status"`
	TransactionsDate   time.Time  `json:"transactions_date"`
	DeliveryStatus     string     `json:"delivery_status" sql:"type:ENUM('pending', 'delivered')"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt          *time.Time `gorm:"column:deleted_at" json:"-"`
}