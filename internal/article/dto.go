package article

import (
	"base-site-api/internal/pagination"
)

// PaginatedArticles struct for articles list and Pagination for api response
type PaginatedArticles struct {
	*pagination.Pagination
	Articles []*Article `json:"articles"`
}
