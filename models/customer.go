package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
}
