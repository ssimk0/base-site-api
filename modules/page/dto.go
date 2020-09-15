package page

import "base-site-api/models"

type PageDetail struct {
	models.Page
	Children []*models.Page `json:"children"`
}
