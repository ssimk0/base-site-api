package models

// Article
type Article struct {
	Model
	Title     string `json:"title" gorm:"not null"`
	Body      string `json:"body" gorm:"not null"`
	Short     string `json:"short" gorm:"not null"`
	Slug      string `json:"slug" gorm:"unique_index;not null"`
	Published bool   `json:"published"`
	Viewed    int    `json:"viewed"`
	UserID    uint   `json:"-" gorm:"not null"`
	User      User   `json:"created_by" gorm:"foreignkey:UserID"`
}
