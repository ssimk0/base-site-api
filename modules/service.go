package modules

type Service struct {
}

func (s *Service) CalculateLimitAndOffset(page int, size int) (int, int) {
	return page - 1*size, size
}
