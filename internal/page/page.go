package page

import (
	"base-site-api/internal/auth"
	"base-site-api/internal/common"
	"database/sql"
)

// PageCategory model
type PageCategory struct {
	common.Model
	Name string `json:"name" gorm:"not null"`
	Slug string `json:"slug" gorm:"unique_index;not null"`
}

// Page model linked to PageCategory
type Page struct {
	common.Model
	Title        string        `json:"title" gorm:"not null"`
	Body         string        `json:"body" gorm:"not null"`
	Slug         string        `json:"slug" gorm:"unique_index;not null"`
	UserID       uint          `json:"-" gorm:"not null"`
	User         auth.User     `json:"created_by" gorm:"foreignkey:UserID"`
	CategoryID   uint          `json:"-" gorm:"not null"`
	ParentPageID sql.NullInt32 `json:"-" grom:"index:idx_parent_page;default:null"`
	ParentPage   *Page         `json:"parent" gorm:"foreignkey:ParentPageID"`
	Category     PageCategory  `json:"page_category" gorm:"foreignkey:CategoryID"`
}

// Service interface for pages
type Service interface {
	FindCategories() ([]*PageCategory, error)
	FindBySlug(slug string) (*PageDetail, error)
	FindAllByCategory(categorySlug string) ([]*Page, error)
	Update(page *Page, id uint) error
	Store(page *Page, categorySlug string, userID uint) (uint, error)
	Delete(id uint, userID uint) error
}

// Repository interface for Page model
type Repository interface {
	FindCategories() ([]*PageCategory, error)
	FindCategoryBySlug(slug string) (*PageCategory, error)
	FindBySlug(slug string) (*Page, []*Page, error)
	FindAllByCategorySlug(categorySlug string) ([]*Page, error)
	Update(page *Page, id uint) error
	Store(page *Page, userID uint) (uint, error)
	Delete(id uint) error
}
