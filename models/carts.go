package models

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	CustomersID uint `json:"customer_id" form:"customer_id"`
	Transactions Transactions
	CartsDetail []CartsDetail
}

type CartsDetail struct {
	gorm.Model
	CartsID    uint `json:"cart_id" form:"cart_id"`
	ProductsID uint `json:"product_id" form:"product_id"`
	Quantity   int  `json:"quantity" form:"quantity"`
}
