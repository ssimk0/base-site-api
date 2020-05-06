package models

import (
	"github.com/jinzhu/gorm"
)

// User model will used for auth
type User struct {
	gorm.Model
	Email           string `gorm:"type:varchar(255);unique_index"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PasswordHash    string `json:"-"`
	Password        string `json:"password" gorm:"-"`
	PasswordConfirm string `json:"password_confirmation" gorm:"-"`
	CanEdit         bool   `json:"is_editor"`
	IsAdmin         bool   `json:"is_admin"`
}
