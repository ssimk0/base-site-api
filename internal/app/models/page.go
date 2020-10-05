package models

import (
	"database/sql"
)

// PageCategory model
type PageCategory struct {
	Model
	Name string `json:"name" gorm:"not null"`
	Slug string `json:"slug" gorm:"unique_index;not null"`
}

// Page model linked to PageCategory
type Page struct {
	Model
	Title        string        `json:"title" gorm:"not null"`
	Body         string        `json:"body" gorm:"not null"`
	Slug         string        `json:"slug" gorm:"unique_index;not null"`
	UserID       uint          `json:"-" gorm:"not null"`
	User         User          `json:"created_by" gorm:"foreignkey:UserID"`
	CategoryID   uint          `json:"-" gorm:"not null"`
	ParentPageID sql.NullInt32 `json:"-" grom:"index:idx_parent_page;default:null"`
	ParentPage   *Page         `json:"parent" gorm:"foreignkey:ParentPageID"`
	Category     PageCategory  `json:"page_category" gorm:"foreignkey:CategoryID"`
}
