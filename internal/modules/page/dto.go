package page

import "base-site-api/internal/app/models"

type PageDetail struct {
	models.Page
	Children []*models.Page `json:"children"`
}
