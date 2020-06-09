package article

import (
	"base-site-api/models"
	"base-site-api/modules"
)

// PaginatedArticles struct for articles list and Pagination for api response
type PaginatedArticles struct {
	*modules.Pagination
	Articles []*models.Article `json:"articless"`
}
