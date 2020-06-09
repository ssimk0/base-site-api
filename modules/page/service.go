package page

import (
	"base-site-api/models"
	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
)

// Service interface for pages
type Service interface {
	FindCategories() ([]*models.PageCategory, error)
	FindBySlug(slug string) (*models.Page, error)
	FindAllByCategory(categorySlug string) ([]*models.Page, error)
	Update(page *models.Page, id uint) error
	Store(page *models.Page, categorySlug string, userID uint) (uint, error)
	Delete(id uint, userID uint) error
}

// NewService return instance of service implementation
func NewService(r Repository) Service {
	return &service{
		r,
	}
}

type service struct {
	repository Repository
}

// FindCategories return all categories
func (s *service) FindCategories() ([]*models.PageCategory, error) {
	return s.repository.FindCategories()
}

// FindBySlug return page by slug
func (s *service) FindBySlug(slug string) (*models.Page, error) {
	return s.repository.FindBySlug(slug)
}

//  FindAllByCategory return pages based on category slug
func (s *service) FindAllByCategory(categorySlug string) ([]*models.Page, error) {
	return s.repository.FindAllByCategorySlug(categorySlug)
}

// Update page
func (s *service) Update(page *models.Page, id uint) error {
	return s.repository.Update(page, id)
}

// Store page
func (s *service) Store(page *models.Page, categorySlug string, userID uint) (uint, error) {
	c, err := s.repository.FindCategoryBySlug(categorySlug)
	if err != nil || c == nil {
		return 0, err
	}
	page.Slug = slug.Make(page.Title)
	page.CategoryID = c.ID
	return s.repository.Store(page, userID)
}

// Delete the page
func (s *service) Delete(id uint, userID uint) error {
	log.Infof("Page with id %d deleted by user with id %d", id, userID)

	return s.repository.Delete(id)
}
