package models

import (
	"time"
)

type Products struct {
	ID           uint          `gorm:"primarykey"`
	Name         string        `json:"name" form:"name"`
	Description  string        `json:"description" form:"description"`
	Price        int           `json:"price" form:"price"`
	Stock        int           `json:"stock" form:"stock"`
	CategoriesID uint          `json:"-" form:"category_id"`
	CartsDetail  []CartsDetail `json:"-"`
	CreatedAt    time.Time     `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time     `gorm:"column:updated_at" json:"-"`
	DeletedAt    *time.Time    `gorm:"column:deleted_at" json:"-"`
}
