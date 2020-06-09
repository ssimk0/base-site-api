package models

// Upload store information about  specific file mainly URL to s3
type Upload struct {
	Model
	File       string         `json:"file" gorm:"not null"`
	Thumbnail  string         `json:"thumbnail"`
	CategoryID uint           `json:"-" gorm:"not null"`
	Category   UploadCategory `json:"category" gorm:"foreignkey:CategoryID"`
}

// UploadCategory will be used to build tree structure of uploads
// also main part of querying specific uploads
type UploadCategory struct {
	Model
	Name    string     `json:"name" gorm:"not null"`
	Slug    string     `json:"slug" gorm:"unique_index;not null"`
	SubPath string     `json:"subpath" gorm:"not null"`
	Type    UploadType `json:"type" gorm:"foreignkey:TypeID"`
	TypeID  uint       `json:"-" gorm:"not null"`
}

// UploadType will be used to enable uploads for specific parts of site
// also used  as querying uploads categories
type UploadType struct {
	Model
	Name string `json:"name" gorm:"not null"`
	Slug string `json:"slug" gorm:"unique_index;not null"`
}
