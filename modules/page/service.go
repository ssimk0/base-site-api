package page

import "base-site-api/models"

type Service interface {
	FindCategories() ([]*models.PageCategory, error)
	FindBySlug(slug string) (*models.Page, error)
	FindAllByCategory(categorySlug string) ([]*models.Page, error)
	Update(page *models.Page, id uint) error
	Store(page *models.Page, userID uint) (uint, error)
	Delete(id uint) error
}

func NewService(r Repository) Service {
	return &service{
		r,
	}
}

type service struct {
	repository Repository
}

func (s *service) FindCategories() ([]*models.PageCategory, error) {
	return s.repository.FindCategories()
}

func (s *service) FindBySlug(slug string) (*models.Page, error) {
	return s.repository.FindBySlug(slug)
}

func (s *service) FindAllByCategory(categorySlug string) ([]*models.Page, error) {
	return s.repository.FindAllByCategorySlug(categorySlug)
}

func (s *service) Update(page *models.Page, id uint) error {
	return s.repository.Update(page, id)
}

func (s *service) Store(page *models.Page, userID uint) (uint, error) {
	return s.repository.Store(page, userID)
}

func (s *service) Delete(id uint) error {
	return s.repository.Delete(id)
}
