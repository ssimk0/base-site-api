package upload

import (
	"base-site-api/internal/models"
	"mime/multipart"
)

// Service interface for upload
type Service interface {
	Store(file *multipart.FileHeader, desc string, categorySlug string, typeSlug string) (*models.Upload, error)
}

// Repository interface of upload
type Repository interface {
	FindCategoriesByType(typeSlug string) ([]*models.UploadCategory, error)
	FindUploadsByCategory(categorySlug string, page int, size int) ([]*models.Upload, int, error)
	Update(desc string, id uint) error
	UpdateCategory(categoryName string, subpath string, thum string, id uint) error
	Find(id uint) (*models.Upload, error)
	FindLatestUploadByCategory(categorySlug string) (*models.Upload, error)
	FindCategory(id uint) (*models.UploadCategory, error)
	FindCategoryBySlug(slug string) (*models.UploadCategory, error)
	FindTypeBySlug(slug string) (*models.UploadType, error)
	Store(upload *models.Upload) (uint, error)
	StoreCategory(categoryName string, subPath string, desc string, typeSlug string) (uint, error)
	Delete(id uint) error
	DeleteCategory(id uint) error
}
