package upload

import (
	"base-site-api/internal/common"
	"mime/multipart"
)

// Upload store information about  specific file mainly URL to s3
type Upload struct {
	common.Model
	File        string         `json:"file" gorm:"not null"`
	Thumbnail   string         `json:"thumbnail"`
	Description string         `json:"description"`
	CategoryID  uint           `json:"-" gorm:"not null"`
	Category    UploadCategory `json:"category" gorm:"foreignkey:CategoryID"`
}

// UploadCategory will be used to build tree structure of upload
// also main part of querying specific upload
type UploadCategory struct {
	common.Model
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description"`
	Slug        string     `json:"slug" gorm:"unique_index;not null"`
	SubPath     string     `json:"subpath" gorm:"not null"`
	Thumbnail   string     `json:"thumbnail"`
	Type        UploadType `json:"type" gorm:"foreignkey:TypeID"`
	TypeID      uint       `json:"-" gorm:"not null"`
}

// UploadType will be used to enable upload for specific parts of site
// also used  as querying upload categories
type UploadType struct {
	common.Model
	Name string `json:"name" gorm:"not null"`
	Slug string `json:"slug" gorm:"unique_index;not null"`
}

// Service interface for upload
type Service interface {
	UploadCategories(typeSlug string) ([]*UploadCategory, error)
	UploadsByCategory(categorySlug string, page int, size int) ([]*Upload, int, error)
	Store(file *multipart.FileHeader, categorySlug string, typeSlug string) (*Upload, error)
	StoreCategory(categoryName string, subPath string, desc string, typeSlug string) (uint, error)
	UpdateCategory(categoryName string, subPath string, id uint) error
	Update(description string, id uint) error
	Delete(id uint) error
	DeleteCategory(id uint) error
}

// Repository interface of upload
type Repository interface {
	FindCategoriesByType(typeSlug string) ([]*UploadCategory, error)
	FindUploadsByCategory(categorySlug string, offset int, limit int) ([]*Upload, int, error)
	Update(desc string, id uint) error
	UpdateCategory(categoryName string, subpath string, thum string, id uint) error
	Find(id uint) (*Upload, error)
	FindCategory(id uint) (*UploadCategory, error)
	FindCategoryBySlug(slug string) (*UploadCategory, error)
	FindTypeBySlug(slug string) (*UploadType, error)
	Store(upload *Upload) (uint, error)
	StoreCategory(category *UploadCategory) (uint, error)
	Delete(id uint) error
	DeleteCategory(id uint) error
}
