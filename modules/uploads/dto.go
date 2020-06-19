package uploads

import (
	"base-site-api/models"
	"base-site-api/modules"
)

// PaginatedUploads struct for uploads list and Pagination for api response
type PaginatedUploads struct {
	*modules.Pagination
	Uploads []*models.Upload `json:"uploads"`
}
