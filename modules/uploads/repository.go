package uploads

import (
	"base-site-api/models"
	"github.com/jinzhu/gorm"
)

// Repository interface of uploads
type Repository interface {
	FindCategoriesByType(typeSlug string) ([]*models.UploadCategory, error)
	FindUploadsByCategory(categorySlug string) ([]*models.Upload, error)
	Update(upload *models.Upload, id uint) error
	UpdateCategory(category *models.UploadCategory, id uint) error
	Store(upload *models.Upload) (uint, error)
	StoreCategory(category *models.UploadCategory) (uint, error)
	Delete(id uint) error
	DeleteCategory(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) FindCategoriesByType(typeSlug string) ([]*models.UploadCategory, error) {
	var t models.UploadType
	var c []*models.UploadCategory

	if err := r.db.Where("slug = ?", typeSlug).Find(&t).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("type_id = ? ", t.ID).Find(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (r *repository) FindUploadsByCategory(categorySlug string) ([]*models.Upload, error) {
	var u []*models.Upload

	var c models.UploadCategory

	if err := r.db.Where("slug = ?", categorySlug).Find(&c).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("category_id = ? ", c.ID).Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repository) Update(upload *models.Upload, id uint) error {
	return nil
}

func (r *repository) UpdateCategory(category *models.UploadCategory, id uint) error {
	return nil
}

func (r *repository) Store(upload *models.Upload) (uint, error) {
	return 0, nil
}

func (r *repository) StoreCategory(category *models.UploadCategory) (uint, error) {
	return 0, nil
}

func (r *repository) Delete(id uint) error {
	return nil
}

func (r *repository) DeleteCategory(id uint) error {
	return nil
}
