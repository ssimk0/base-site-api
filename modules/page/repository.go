package page

import "github.com/jinzhu/gorm"

// TODO: define new repository interface for module page
type Repository struct {
	db *gorm.DB
}
