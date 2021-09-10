package models

import "gorm.io/gorm"

type Deliveries struct {
	gorm.Model
	TransactionsID uint
	DeliveryStatus string
}
