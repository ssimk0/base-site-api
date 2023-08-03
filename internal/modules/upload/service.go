package upload

import (
	"base-site-api/internal/log"
	"base-site-api/internal/models"
	"base-site-api/internal/pagination"
	"base-site-api/internal/storage"
	"fmt"
	"mime/multipart"

	"github.com/gosimple/slug"
)

func NewService(r Repository) Service {
	return &service{
		repository: r,
		store:      storage.Instance(),
	}
}

type service struct {
	pagination.Service
	repository Repository
	store      storage.Storage
}

// UploadCategories by type slug
func (s *service) UploadCategories(typeSlug string) ([]*models.UploadCategory, error) {
	return s.repository.FindCategoriesByType(typeSlug)
}

// Uploads by category slug
func (s *service) UploadsByCategory(categorySlug string, page int, size int) ([]*models.Upload, int, error) {
	l, o := s.CalculateLimitAndOffset(page, size)
	return s.repository.FindUploadsByCategory(categorySlug, l, o)
}

// Store upload the storage and save the row to db with all information about the storage itself
func (s *service) Store(file *multipart.FileHeader, desc string, categorySlug string, typeSlug string) (*models.Upload, error) {
	category, err := s.repository.FindCategoryBySlug(categorySlug)
	if err != nil {
		return nil, err
	}

	t, err := s.repository.FindTypeBySlug(typeSlug)
	if err != nil {
		return nil, err
	}

	f, err := s.store.Store(file, fmt.Sprintf("%s/%s", t.Slug, category.SubPath))

	if err != nil {
		return nil, err
	}

	c, err := s.repository.FindCategoryBySlug(categorySlug)

	if err != nil {
		return nil, err
	}

	u := models.Upload{
		File:        f.URL,
		Thumbnail:   f.URLSmall,
		Description: desc,
		CategoryID:  c.ID,
	}

	_, err = s.repository.Store(&u)

	if err != nil {
		return nil, err
	}

	if c.Thumbnail == "" {
		err := s.repository.UpdateCategory(c.Name, c.SubPath, u.File, c.ID)
		if err != nil {
			log.Errorf("Error while update the category %s", err.Error())
		}
	}
	return &u, nil
}

func (s *service) StoreCategory(categoryName string, subPath string, desc string, typeSlug string) (uint, error) {
	t, err := s.repository.FindTypeBySlug(typeSlug)
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

	return s.repository.StoreCategory(&c)
}

// UpdateCategory update the category it self and later also the s3 bucket
func (s *service) UpdateCategory(categoryName string, subPath string, id uint) error {
	c, err := s.repository.FindCategory(id)
	if err != nil {
		return err
	}

	if categoryName == "" {
		categoryName = c.Name
	}

	if subPath == "" {
		subPath = c.SubPath
	}

	return s.repository.UpdateCategory(categoryName, subPath, "", id)
}

func (s *service) Update(desc string, id uint) error {
	return s.repository.Update(desc, id)
}

func (s *service) Delete(id uint) error {
	return s.repository.Delete(id)
}

func (s *service) DeleteCategory(id uint) error {
	return s.repository.DeleteCategory(id)
}

func (s *service) LatestUpload(categorySlug string) (*models.Upload, error) {
	return s.repository.FindLatestUploadByCategory(categorySlug)
}

func (s *service) Upload(id uint) (*models.Upload, error) {
	return s.repository.Find(id)
}
