package article

import (
	"base-site-api/models"
)

// Service interface for Article model
type Service interface {
	Find(slug string) (*models.Article, error)
	FindAll(sort string, page int, size int) ([]*models.Article, int, error)
	Update(article *models.Article, id uint) error
	Store(article *models.Article, userID uint) (*models.Article, error)
	Delete(id uint, userID uint) error
}
