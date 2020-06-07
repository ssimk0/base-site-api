package page

import (
	"base-site-api/models"
	"github.com/jinzhu/gorm"
)

//Repository interface for Page model
type Repository interface {
	FindCategories() ([]*models.PageCategory, error)
	FindCategoryBySlug(slug string) (*models.PageCategory, error)
	FindBySlug(slug string) (*models.Page, error)
	FindAll(categorySlug string) ([]*models.Page, error)
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

	if err := r.db.Find(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (r *repository) FindCategoryBySlug(slug string) (*models.PageCategory, error) {
	var c models.PageCategory

	if err := r.db.Model(&models.PageCategory{}).Where("slug = ?", slug).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *repository) FindAll(categorySlug string) ([]*models.Page, error) {
	var pages []*models.Page
	category, err := r.FindCategoryBySlug(categorySlug)
	if err != nil {
		return nil, err
	}

	if err := r.db.Set("gorm:auto_preload", true).Model(&models.Page{}).Where("category_id = ?", category.ID).Find(&pages).Error; err != nil {
		return nil, err
	}

	return pages, nil
}

func (r *repository) FindBySlug(slug string) (*models.Page, error) {
	var c models.Page

	if err := r.db.Model(c).Where("slug = ?", slug).Find(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *repository) Find(id uint) (*models.Page, error) {
	page := models.Page{}

	if err := r.db.First(&page, id).Error; err != nil {
		return nil, err
	}

	return &page, nil
}

func (r *repository) Update(page *models.Page, id uint) error {
	p, err := r.Find(id)
	if err != nil {
		return err
	}

	if page.Title != "" {
		p.Title = page.Title
	}

	if page.Body != "" {
		p.Body = page.Body
	}

	if page.Slug != "" {
		p.Slug = page.Slug
	}

	return r.db.Save(p).Error
}

func (r *repository) Store(page *models.Page, userID uint) (uint, error) {
	page.UserID = userID
	if err := r.db.Create(page).Error; err != nil {
		return 0, err
	}

	return page.ID, nil
}

func (r *repository) Delete(id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Page{}).Error; err != nil {
		return err
	}

	return nil
}
