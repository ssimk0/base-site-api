package upload

import (
	"base-site-api/internal/file"
	"base-site-api/internal/log"
	"base-site-api/internal/pagination"
	"fmt"
	"mime/multipart"

	"github.com/gosimple/slug"
)

func NewService(r Repository) Service {
	return &service{
		repository: r,
		s3:         file.NewS3(),
	}
}

type service struct {
	pagination.Service
	repository Repository
	s3         *file.S3Storage
}

// UploadCategories by type slug
func (s *service) UploadCategories(typeSlug string) ([]*UploadCategory, error) {
	return s.repository.FindCategoriesByType(typeSlug)
}

// Uploads by category slug
func (s *service) UploadsByCategory(categorySlug string, page int, size int) ([]*Upload, int, error) {
	l, o := s.CalculateLimitAndOffset(page, size)
	return s.repository.FindUploadsByCategory(categorySlug, l, o)
}

// Store upload the file and save the row to db with all information about the file itself
func (s *service) Store(file *multipart.FileHeader, categorySlug string, typeSlug string) (*Upload, error) {
	category, err := s.repository.FindCategoryBySlug(categorySlug)
	if err != nil {
		return nil, err
	}

	t, err := s.repository.FindTypeBySlug(typeSlug)
	if err != nil {
		return nil, err
	}

	f, err := s.s3.Store(file, fmt.Sprintf("%s/%s", t.Slug, category.SubPath))

	if err != nil {
		return nil, err
	}

	c, err := s.repository.FindCategoryBySlug(categorySlug)

	if err != nil {
		return nil, err
	}

	u := Upload{
		File:       f.URL,
		Thumbnail:  f.URLSmall,
		CategoryID: c.ID,
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

	c := UploadCategory{
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
