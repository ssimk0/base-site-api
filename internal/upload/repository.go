package upload

import (
	"github.com/jinzhu/gorm"
)

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
func (r *repository) FindCategoriesByType(typeSlug string) ([]*UploadCategory, error) {
	var t UploadType
	var c []*UploadCategory

	if err := r.db.Where("slug = ?", typeSlug).Find(&t).Error; err != nil {
		return nil, err
	}

	if err := r.db.Set("gorm:auto_preload", true).Where("type_id = ? ", t.ID).Find(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// FindUploadsByCategory return all upload for the category
func (r *repository) FindUploadsByCategory(categorySlug string, offset int, limit int) ([]*Upload, int, error) {
	var u []*Upload
	var count int

	var c UploadCategory
	if err := r.db.Where("slug = ?", categorySlug).Find(&c).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Set("gorm:auto_preload", true).Where("category_id = ? ", c.ID).Offset(offset).Limit(limit).Find(&u).Count(&count).Error; err != nil {
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
func (r *repository) Find(id uint) (*Upload, error) {
	upload := Upload{}

	if err := r.db.First(&upload, id).Error; err != nil {
		return nil, err
	}

	return &upload, nil
}

// FindCategory by id
func (r *repository) FindCategory(id uint) (*UploadCategory, error) {
	category := UploadCategory{}

	if err := r.db.Set("gorm:auto_preload", true).First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// FindCategoryBySlug used for create upload
func (r *repository) FindCategoryBySlug(slug string) (*UploadCategory, error) {
	category := UploadCategory{}

	if err := r.db.Set("gorm:auto_preload", true).Where("slug = ?", slug).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *repository) FindTypeBySlug(slug string) (*UploadType, error) {
	t := UploadType{}

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
	u.SubPath = subpath
	u.Name = categoryName
	if thum != "" {
		u.Thumbnail = thum
	}
	return r.db.Save(&u).Error
}

// Store upload
func (r *repository) Store(upload *Upload) (uint, error) {
	if err := r.db.Create(upload).Error; err != nil {
		return 0, err
	}

	return upload.ID, nil
}

// StoreCategory upload
func (r *repository) StoreCategory(category *UploadCategory) (uint, error) {
	if err := r.db.Create(category).Error; err != nil {
		return 0, err
	}

	return category.ID, nil
}

// Delete upload from DB
func (r *repository) Delete(id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&Upload{}).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCategory upload
func (r *repository) DeleteCategory(id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&UploadCategory{}).Error; err != nil {
		return err
	}

	return nil
}
