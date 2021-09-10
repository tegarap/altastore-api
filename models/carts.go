package models

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	CustomersID uint `json:"customer_id" form:"customer_id"`
	//Transactions Transactions
	CartsDetail []CartsDetail
}

type CartsDetail struct {
	gorm.Model
	CartsID  uint
	Products []*Products `gorm:"many2many:carts_products;" json:"product_id" form:"product_id"`
	Quantity int         `json:"quantity" form:"quantity"`
}
