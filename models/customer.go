package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"-" form:"password"`
	Phone        int    `json:"phone" form:"phone"`
	Address      string `json:"address" form:"address"`
	Gender       string `json:"gender" sql:"type:ENUM('male', 'female')"`
	Carts        []Carts
	Transactions []Transactions
}

type CustomerResponse struct {
	ID       uint
	Name     string
	Email    string
	Phone    int
	Address  string
	Gender   string
}