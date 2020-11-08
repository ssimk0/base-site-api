package page

import (
	"base-site-api/internal/app/models"
)

// Service interface for pages
type Service interface {
	FindCategories() ([]*models.PageCategory, error)
	FindBySlug(slug string) (*PageDetail, error)
	FindAllByCategory(categorySlug string) ([]*PageDetail, error)
	Update(page *models.Page, id uint) error
	Store(page *models.Page, categorySlug string, userID uint) (uint, error)
	Delete(id uint, userID uint) error
}

// Repository interface for Page model
type Repository interface {
	FindCategories() ([]*models.PageCategory, error)
	FindCategoryBySlug(slug string) (*models.PageCategory, error)
	FindBySlug(slug string) (*models.Page, []*models.Page, error)
	FindAllByCategorySlug(categorySlug string) ([]*PageDetail, error)
	Update(page *models.Page, id uint) error
	Store(page *models.Page, userID uint) (uint, error)
	Delete(id uint) error
}
