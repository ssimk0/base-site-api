package uploads

import (
	"base-site-api/models"
	"mime/multipart"
)

// Service interface for uploads
type Service interface {
	UploadCategories(typeSlug string) ([]*models.PageCategory, error)
	UploadsByCategory(categorySlug string) ([]*models.Upload, error)
	Store(file *multipart.FileHeader, filename string, categorySlug string) (uint, error)
	StoreCategory(categoryName string, subPath string, typeSlug string) (uint, error)
	UpdateCategory(categoryName string, subPath string, typeSlug string) error
	Update(description string, id uint) error
	Delete(id uint) error
	DeleteCategory(id uint) error
}
