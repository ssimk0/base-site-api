package upload

import (
	"base-site-api/internal/log"
	"base-site-api/internal/models"
	"base-site-api/internal/pagination"
	"base-site-api/internal/storage"
	"fmt"
	"mime/multipart"
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
