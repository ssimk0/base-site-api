package upload

import (
	"base-site-api/internal/models"
	"base-site-api/internal/pagination"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

type repository struct {
	pagination.Service
	db *gorm.DB
}

// NewRepository for the upload module
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		pagination.Service{},
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

	if err := r.db.Set("gorm:auto_preload", true).Where("type_id = ? ", t.ID).Order("created_at desc").Find(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// FindUploadsByCategory return all upload for the category
func (r *repository) FindUploadsByCategory(categorySlug string, page int, size int) ([]*models.Upload, int, error) {
	offset, limit := r.CalculateLimitAndOffset(page, size)
	var u []*models.Upload
	var count int

	var c models.UploadCategory
	if err := r.db.Where("slug = ?", categorySlug).Find(&c).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Set("gorm:auto_preload", true).Model(&models.Upload{}).Order("created_at desc").Where("category_id = ? ", c.ID).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Set("gorm:auto_preload", true).Where("category_id = ? ", c.ID).Order("created_at desc").Offset(offset).Limit(limit).Find(&u).Error; err != nil {
		return nil, 0, err
	}

	return u, count, nil
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

	if err := r.db.Set("gorm:auto_preload", true).First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// FindCategoryBySlug used for create upload
func (r *repository) FindCategoryBySlug(slug string) (*models.UploadCategory, error) {
	category := models.UploadCategory{}

	if err := r.db.Set("gorm:auto_preload", true).Where("slug = ?", slug).First(&category).Error; err != nil {
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

// UpdateCategory only subpath and name now
func (r *repository) UpdateCategory(categoryName string, subpath string, thum string, id uint) error {
	u, err := r.FindCategory(id)
	if err != nil {
		return err
	}

	if categoryName != "" {
		u.Name = categoryName
	}
	if subpath != "" {
		u.SubPath = subpath
	}

	if thum != "" {
		u.Thumbnail = thum
	}

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
func (r *repository) StoreCategory(categoryName string, subPath string, desc string, typeSlug string) (uint, error) {
	t, err := r.FindTypeBySlug(typeSlug)
	if err != nil {
		return 0, err
	}

	c := models.UploadCategory{
		Name:        categoryName,
		SubPath:     subPath,
		Description: desc,
		TypeID:      t.ID,
		Slug:        slug.Make(categoryName),
	}

	if err := r.db.Create(&c).Error; err != nil {
		return 0, err
	}

	return c.ID, nil
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

func (r *repository) FindLatestUploadByCategory(categorySlug string) (*models.Upload, error) {
	c, err := r.FindCategoryBySlug(categorySlug)
	upload := models.Upload{}

	if err != nil {
		return &upload, err
	}

	if err := r.db.Set("gorm:auto_preload", true).Where("category_id = ?", c.ID).Order("created_at desc").First(&upload).Error; err != nil {
		return nil, err
	}

	return &upload, nil
}
