package article

import (
	"base-site-api/internal/app/models"
)

//Repository interface for Article model
type Repository interface {
	Find(id uint) (*models.Article, error)
	FindBySlug(slug string) (*models.Article, error)
	FindAll(order string, offset int, limit int) ([]*models.Article, int, error)
	Update(article *models.Article, id uint) error
	Store(article *models.Article, userID uint) (uint, error)
	Delete(id uint) error
}

// Service interface for Article model
type Service interface {
	Find(slug string) (*models.Article, error)
	FindAll(sort string, page int, size int) ([]*models.Article, int, error)
	Update(article *models.Article, id uint) error
	Store(article *models.Article, userID uint) (*models.Article, error)
	Delete(id uint, userID uint) error
}
