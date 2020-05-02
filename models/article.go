package models

import (
	"github.com/jinzhu/gorm"
)

// Article
type Article struct {
	gorm.Model
	Title     string `json:"title"`
	Body      string `json:"body"`
	Short     string `json:"short"`
	Slug      string `json:"slug" gorm:"unique_index;not null"`
	Published bool   `json:"published"`
	Viewed    int    `json:"viewed"`
	CreatedBy User   `json:"created_by"`
}
