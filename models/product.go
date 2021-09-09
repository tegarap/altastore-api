package models

import "gorm.io/gorm"

type ProductsSchema struct {
	gorm.Model
	Category    Categories
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}
