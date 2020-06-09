package uploads

import "base-site-api/models"

// Repository interface of uploads
type Repository interface {
	FindCategoriesByType(typeSlug string) ([]*models.UploadCategory, error)
	FindUploadsByCategory(categorySlug uint) ([]*models.Upload, error)
	Update(upload *models.Upload, id uint) error
	UpdateCategory(category *models.UploadCategory, id uint) error
	Store(upload *models.Upload) (uint, error)
	StoreCategory(category *models.UploadCategory) (uint, error)
	Delete(id uint) error
	DeleteCategory(id uint) error
}
