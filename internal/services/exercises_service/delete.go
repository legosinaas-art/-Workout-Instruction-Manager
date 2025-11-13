package exercises_service

import (
	"context"
	"fmt"
)

func (s *Service) Delete(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete exercises_repo: %w", err)
	}
	return nil
}
