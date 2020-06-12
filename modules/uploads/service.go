package uploads

import (
	"base-site-api/models"
	"base-site-api/storage"
	"mime/multipart"
)

// Service interface for uploads
type Service interface {
	UploadCategories(typeSlug string) ([]*models.UploadCategory, error)
	UploadsByCategory(categorySlug string) ([]*models.Upload, error)
	Store(file *multipart.FileHeader, filename string, categorySlug string) (uint, error)
	StoreCategory(categoryName string, subPath string, typeSlug string) (uint, error)
	UpdateCategory(categoryName string, subPath string, typeSlug string) error
	Update(description string, id uint) error
	Delete(id uint) error
	DeleteCategory(id uint) error
}

type service struct {
	r  Repository
	s3 storage.S3Storage
}

// UploadCategories by type slug
func (s *service) UploadCategories(typeSlug string) ([]*models.UploadCategory, error) {
	return s.UploadCategories(typeSlug)
}

// Uploads by category slug
func (s *service) UploadsByCategory(categorySlug string) ([]*models.Upload, error) {
	return s.UploadsByCategory(categorySlug)
}

// Store upload the file and save the row to db with all information about the file itself
func (s *service) Store(file *multipart.FileHeader, filename string, categorySlug string) (*storage.UploadFile, error) {
	f, err := s.s3.Store(file, filename)

	if err != nil {
		return nil, err
	}

	c, err := s.r.FindCategoryBySlug(categorySlug)

	if err != nil {
		return nil, err
	}

	u := models.Upload{
		File:       f.URL,
		Thumbnail:  f.URLSmall,
		CategoryID: c.ID,
	}

	_, err = s.r.Store(&u)

	if err != nil {
		return nil, err
	}

	return f, nil
}
