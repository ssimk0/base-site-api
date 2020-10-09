package storage

import (
	"base-site-api/internal/app/config"
	"mime/multipart"
)

// UploadFile is return value from Store function from Storage
type UploadFile struct {
	URL      string `json:"url"`
	URLSmall string `json:"url-small"`
	IsImage  bool   `json:"-"`
}

type Storage interface {
	Store(file *multipart.FileHeader, fileName string) (*UploadFile, error)
}

var s Storage

func Initialize(c *config.StorageConfiguration) {
	if c.Type == "s3" {
		s = NewS3(c)
	}
}

func Instance() Storage {
	return s
}
