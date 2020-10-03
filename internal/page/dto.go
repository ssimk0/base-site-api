package page

type PageDetail struct {
	Page
	Children []*Page `json:"children"`
}
