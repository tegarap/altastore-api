package models

import "time"

type Admins struct {
	ID        uint       `gorm:"primarykey"`
	Name      string     `gorm:"name" json:"name" form:"name"`
	Email     string     `gorm:"email;unique;not null" json:"email" form:"email"`
	Password  string     `gorm:"password" json:"password" form:"password"`
	Token     string     `gorm:"token" json:"token" form:"token"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}
