package models

import "time"

// User model will used for auth
type User struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	DeletedAt    *time.Time `json:"-" sql:"index"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	Email        string     `json:"email" gorm:"type:varchar(255);unique_index;not null"`
	FirstName    string     `json:"first_name" gorm:"not null"`
	LastName     string     `json:"last_name" gorm:"not null"`
	PasswordHash string     `json:"-" gorm:"not null"`
	CanEdit      bool       `json:"-"`
	IsAdmin      bool       `json:"-"`
}
