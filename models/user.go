package models

import "time"

// User model will used for auth
type User struct {
	ID              uint       `json:"id" gorm:"primary_key"`
	DeletedAt       *time.Time `json:"-" sql:"index"`
	CreatedAt       time.Time  `json:"-"`
	UpdatedAt       time.Time  `json:"-"`
	Email           string     `json:"email" gorm:"type:varchar(255);unique_index"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	PasswordHash    string     `json:"-"`
	Password        string     `json:"password" gorm:"-"`
	PasswordConfirm string     `json:"password_confirmation" gorm:"-"`
	CanEdit         bool       `json:"is_editor"`
	IsAdmin         bool       `json:"is_admin"`
}
