package models

import (
	"time"
)

type Customers struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    int    `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" sql:"type:ENUM('male', 'female')"`
	Token    string `json:"token"`
	Carts    []Carts
	//Transactions []Transactions
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}
