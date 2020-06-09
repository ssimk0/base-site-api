package responses

// SuccessResponse json struct
type SuccessResponse struct {
	Success bool `json:"success"`
	ID      uint `json:"id"`
}

// ErrorResponse json struct
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
