package exercises_service

import (
	"context"
	"errors"
	"fmt"

	"example.com/m/v2/internal/dto"
	"example.com/m/v2/internal/repository/repo_error"
)

func (s *Service) GetById(ctx context.Context, id int) (dto.ExercisesDto, error) {
	exercise, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, repo_error.ErrNotFound) {
			return dto.ExercisesDto{}, repo_error.ErrNotFound
		}
		return dto.ExercisesDto{}, fmt.Errorf("failed to show exercise: %w", err)
	}
	return exercise.ToDto(), nil
}
