package page

import (
	"base-site-api/internal/app/models"
	"base-site-api/internal/log"
	"database/sql"
	"github.com/gosimple/slug"
)

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
func (s *service) FindBySlug(slug string) (*PageDetail, error) {
	page, children, err := s.repository.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	return &PageDetail{
		*page,
		children,
	}, err
}

//  FindAllByCategory return pages based on category slug
func (s *service) FindAllByCategory(categorySlug string) ([]*PageDetail, error) {
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

	if page.ParentPage.ID != 0 {
		id := sql.NullInt32{}
		err = id.Scan(page.ParentPage.ID)
		if err != nil {
			return 0, err
		}

		page.ParentPageID = id
		page.ParentPage = nil
	}
	return s.repository.Store(page, userID)
}

// Delete the page
func (s *service) Delete(id uint, userID uint) error {
	log.Infof("Page with id %d deleted by user with id %d", id, userID)

	return s.repository.Delete(id)
}
