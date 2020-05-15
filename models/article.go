package models

// Article
type Article struct {
	Model
	Title     string `json:"title"`
	Body      string `json:"body"`
	Short     string `json:"short"`
	Slug      string `json:"slug" gorm:"unique_index;not null"`
	Published bool   `json:"published"`
	Viewed    int    `json:"viewed"`
	UserID    uint    `json:"-"`
	User      User   `json:"created_by" gorm:"foreignkey:UserID"`
}
