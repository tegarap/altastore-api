package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Description  string `json:"description" form:"description"`
	Price        int    `json:"price" form:"price"`
	Stock        int    `json:"stock" form:"stock"`
	CategoriesID uint   `json:"category_id" form:"category_id"`
	CartsDetail  []CartsDetail
}

type OutputProduct struct {
	ID          uint
	Name        string
	Description string
	Price       int
	Stock       int
}
