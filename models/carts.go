package models

import (
	"time"
)

type Carts struct {
	ID           uint `gorm:"primarykey"`
	CustomersID  uint `json:"customer_id" form:"customer_id"`
	Transactions Transactions
	CartsDetail  []CartsDetail
	CreatedAt    time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"-"`
}

type CartsDetail struct {
	ID         uint       `gorm:"primarykey"`
	CartsID    uint       `json:"cart_id" form:"cart_id"`
	ProductsID uint       `json:"product_id" form:"product_id"`
	Quantity   int        `json:"quantity" form:"quantity"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"-"`
}
