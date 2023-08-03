package models

import "base-site-api/internal/pagination"

type Article struct {
	Model
	Title     string `json:"title" gorm:"not null"`
	Image     string `json:"image"`
	Body      string `json:"body" gorm:"not null"`
	Short     string `json:"short" gorm:"not null"`
	Slug      string `json:"slug" gorm:"unique_index;not null"`
	Published bool   `json:"published"`
	Viewed    int    `json:"viewed"`
	UserID    uint   `json:"-" gorm:"not null"`
	User      User   `json:"created_by" gorm:"foreignkey:UserID"`
}

// PaginatedArticles struct for articles list and Pagination for api response
type PaginatedArticles struct {
	*pagination.Pagination
	Articles []*Article `json:"articles"`
}
