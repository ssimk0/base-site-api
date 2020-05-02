package responses

type SuccessResponse struct {
	Success bool `json:"success"`
	Id      uint `json:"id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
