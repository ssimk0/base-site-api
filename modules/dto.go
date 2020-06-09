package modules

// Pagination wrap params for all pagination
type Pagination struct {
	Page       int     `json:"page"`
	PageSize   int     `json:"page_size"`
	Total      int     `json:"total"`
	TotalPages float64 `json:"total_pages"`
}
