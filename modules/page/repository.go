package page

import (
	"base-site-api/models"
	"github.com/jinzhu/gorm"
)

//Repository interface for Page model
type Repository interface {
	FindCategories() ([]*models.PageCategory, error)
	FindBySlug(slug string) (*models.Page, error)
	FindAll(categorySlug string) ([]*models.Page, int, error)
	Update(page *models.Page, id uint) error
	Store(page *models.Page, userID uint) (uint, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository return instance of repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) FindCategories() ([]*models.PageCategory, error) {
	var c []*models.PageCategory

	if err := r.db.Model(c).Find(c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (r *repository) FindAll(categorySlug string) ([]*models.Page, int, error) {
	return nil, 0, nil
}

func (r *repository) FindBySlug(slug string) (*models.Page, error) {
	return nil, nil
}

func (r *repository) Update(page *models.Page, id uint) error {
	return nil
}

func (r *repository) Store(page *models.Page, userID uint) (uint, error) {
	return uint(0), nil
}

func (r *repository) Delete(id uint) error {
	return nil
}
