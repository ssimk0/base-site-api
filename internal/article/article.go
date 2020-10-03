package article

import (
	"base-site-api/internal/auth"
	"base-site-api/internal/common"
)

type Article struct {
	common.Model
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body" gorm:"not null"`
	Short     string    `json:"short" gorm:"not null"`
	Slug      string    `json:"slug" gorm:"unique_index;not null"`
	Published bool      `json:"published"`
	Viewed    int       `json:"viewed"`
	UserID    uint      `json:"-" gorm:"not null"`
	User      auth.User `json:"created_by" gorm:"foreignkey:UserID"`
}

//Repository interface for Article model
type Repository interface {
	Find(id uint) (*Article, error)
	FindBySlug(slug string) (*Article, error)
	FindAll(order string, offset int, limit int) ([]*Article, int, error)
	Update(article *Article, id uint) error
	Store(article *Article, userID uint) (uint, error)
	Delete(id uint) error
}

// Service interface for Article model
type Service interface {
	Find(slug string) (*Article, error)
	FindAll(sort string, page int, size int) ([]*Article, int, error)
	Update(article *Article, id uint) error
	Store(article *Article, userID uint) (*Article, error)
	Delete(id uint, userID uint) error
}
