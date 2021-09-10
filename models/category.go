package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Name     string `json:"category" form:"category"`
	Products []Products
}
