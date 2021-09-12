package models

import (
	"time"
)

type Admins struct {
	ID        uint       `gorm:"primarykey"`
	Username  string     `gorm:"unique;not null" json:"username"`
	Name      string     `gorm:"name" json:"name"`
	Email     string     `gorm:"email" json:"email"`
	Password  string     `gorm:"password" json:"password"`
	Token     string     `gorm:"token" json:"token"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}
