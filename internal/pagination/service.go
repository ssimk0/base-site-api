package pagination

// Service wrap common logic for service
type Service struct {
}

// CalculateLimitAndOffset help with calculating offset for pagination
func (s *Service) CalculateLimitAndOffset(page int, size int) (int, int) {
	return (page - 1) * size, size
}
