package uploads

import (
	"base-site-api/models"
	"github.com/jinzhu/gorm"
)

// Repository interface of uploads
type Repository interface {
	FindCategoriesByType(typeSlug string) ([]*models.UploadCategory, error)
	FindUploadsByCategory(categorySlug string) ([]*models.Upload, error)
	Update(desc string, id uint) error
	UpdateCategory(subpath string, id uint) error
	Find(id uint) (*models.Upload, error)
	FindCategory(id uint) (*models.UploadCategory, error)
	FindCategoryBySlug(slug string) (*models.UploadCategory, error)
	FindTypeBySlug(slug string) (*models.UploadType, error)
	Store(upload *models.Upload) (uint, error)
	StoreCategory(category *models.UploadCategory) (uint, error)
	Delete(id uint) error
	DeleteCategory(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository for the upload module
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

// FindCategoriesByType return all categories for upload type
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

// FindUploadsByCategory return all uploads for the category
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

// Update for the upload
func (r *repository) Update(desc string, id uint) error {
	u, err := r.Find(id)
	if err != nil {
		return err
	}
	u.Description = desc
	return r.db.Save(&u).Error
}

// Find upload by id
func (r *repository) Find(id uint) (*models.Upload, error) {
	upload := models.Upload{}

	if err := r.db.First(&upload, id).Error; err != nil {
		return nil, err
	}

	return &upload, nil
}

// FindCategory by id
func (r *repository) FindCategory(id uint) (*models.UploadCategory, error) {
	category := models.UploadCategory{}

	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// FindCategoryBySlug used for create upload
func (r *repository) FindCategoryBySlug(slug string) (*models.UploadCategory, error) {
	category := models.UploadCategory{}

	if err := r.db.Where("slug = ?", slug).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *repository) FindTypeBySlug(slug string) (*models.UploadType, error) {
	t := models.UploadType{}

	if err := r.db.Where("slug = ?", slug).First(&t).Error; err != nil {
		return nil, err
	}

	return &t, nil
}

// UpdateCategory only subpath now
func (r *repository) UpdateCategory(subpath string, id uint) error {
	u, err := r.FindCategory(id)
	if err != nil {
		return err
	}
	u.SubPath = subpath
	return r.db.Save(&u).Error
}

// Store upload
func (r *repository) Store(upload *models.Upload) (uint, error) {
	if err := r.db.Create(upload).Error; err != nil {
		return 0, err
	}

	return upload.ID, nil
}

// StoreCategory upload
func (r *repository) StoreCategory(category *models.UploadCategory) (uint, error) {
	if err := r.db.Create(category).Error; err != nil {
		return 0, err
	}

	return category.ID, nil
}

// Delete upload from DB
func (r *repository) Delete(id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Upload{}).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCategory upload
func (r *repository) DeleteCategory(id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&models.UploadCategory{}).Error; err != nil {
		return err
	}

	return nil
}
