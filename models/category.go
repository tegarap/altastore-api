package models

import (
	"time"
)

type Categories struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `json:"category_name" form:"category_name"`
	Products  []Products
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}
