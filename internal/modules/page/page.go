package page

import (
	"base-site-api/internal/app/models"
)

// Repository interface for Page model
type Repository interface {
	FindCategories() ([]*models.PageCategory, error)
	FindCategoryBySlug(slug string) (*models.PageCategory, error)
	FindBySlug(slug string) (*models.Page, []*models.Page, error)
	FindAllByCategorySlug(categorySlug string) ([]*PageDetail, error)
	Update(page *models.Page, id uint) error
	Store(page *models.Page, categorySlug string, userID uint) (uint, error)
	Delete(id uint, userID uint) error
}
