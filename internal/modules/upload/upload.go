package upload

import (
	"base-site-api/internal/app/models"
	"mime/multipart"
)

// Service interface for upload
type Service interface {
	UploadCategories(typeSlug string) ([]*models.UploadCategory, error)
	LatestUpload(categorySlug string) (*models.Upload, error)
	UploadsByCategory(categorySlug string, page int, size int) ([]*models.Upload, int, error)
	Store(file *multipart.FileHeader, categorySlug string, typeSlug string) (*models.Upload, error)
	StoreCategory(categoryName string, subPath string, desc string, typeSlug string) (uint, error)
	UpdateCategory(categoryName string, subPath string, id uint) error
	Update(description string, id uint) error
	Delete(id uint) error
	DeleteCategory(id uint) error
}

// Repository interface of upload
type Repository interface {
	FindCategoriesByType(typeSlug string) ([]*models.UploadCategory, error)
	FindUploadsByCategory(categorySlug string, offset int, limit int) ([]*models.Upload, int, error)
	Update(desc string, id uint) error
	UpdateCategory(categoryName string, subpath string, thum string, id uint) error
	Find(id uint) (*models.Upload, error)
	FindLatestUploadByCategory(categorySlug string) (*models.Upload, error)
	FindCategory(id uint) (*models.UploadCategory, error)
	FindCategoryBySlug(slug string) (*models.UploadCategory, error)
	FindTypeBySlug(slug string) (*models.UploadType, error)
	Store(upload *models.Upload) (uint, error)
	StoreCategory(category *models.UploadCategory) (uint, error)
	Delete(id uint) error
	DeleteCategory(id uint) error
}
