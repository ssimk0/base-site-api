package upload

import (
	"base-site-api/internal/models"
	"base-site-api/internal/pagination"
)

// PaginatedUploads struct for upload list and Pagination for api response
type PaginatedUploads struct {
	*pagination.Pagination
	Uploads []*models.Upload `json:"upload"`
}
