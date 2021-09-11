package models

import (
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name     string `json:"category_name" form:"category_name"`
	Products []Products
}
