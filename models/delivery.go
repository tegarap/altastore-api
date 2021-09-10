package models

import "gorm.io/gorm"

type Deliveries struct {
	gorm.Model
	TransactionsID uint   `json:"transactions_id"`
	DeliveryStatus string `json:"delivery_status"`
}
