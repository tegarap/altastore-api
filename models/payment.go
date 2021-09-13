package models

import (
	"time"
)

type Payments struct {
	ID           uint           `gorm:"primarykey"`
	PaymentName  string         `json:"payment_name"`
	Transactions []Transactions `json:"-"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"-"`
	DeletedAt    *time.Time     `gorm:"column:deleted_at" json:"-"`
}
