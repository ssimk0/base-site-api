package page

import "base-site-api/models"

type Service interface {
	FindCategories() ([]*models.PageCategory, error)
	FindBySlug() (*models.Page, error)
	FindAllByCategory(categorySlug string) ([]*models.Page, error)
	Update(page *models.Page, id uint) error
	Store(page *models.Page, userID uint) (uint, error)
	Delete(id uint) error
}

// TODO: define new service for module page
type service struct {
	repository Repository
}
