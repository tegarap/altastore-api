package models

import "gorm.io/gorm"

type Admins struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminResponse struct {
	ID       uint
	Name     string
	Email    string
}