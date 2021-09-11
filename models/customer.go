package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    int    `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender"`
	Carts    []Carts
	//Transactions []Transactions
}
