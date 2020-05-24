package article

import (
	"base-site-api/models"
	"base-site-api/modules"
)

type PaginatedArticles struct {
	*modules.Pagination
	Articles []*models.Article `json:"articless"`
}
