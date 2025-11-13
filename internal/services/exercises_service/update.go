package exercises_service

import (
	"context"
	"fmt"

	"example.com/m/v2/internal/dto"
)

func (s *Service) Update(ctx context.Context, req dto.ExercisesDto) (dto.ExercisesDto, error) {
	exercise, err := s.repo.GetByID(ctx, req.Id)
	if err != nil {
		return dto.ExercisesDto{}, fmt.Errorf("not found id %w", err)
	}
	exercise.Name = req.Name
	exercise.Notes = req.Notes
	err = s.repo.Update(ctx, exercise)
	if err != nil {
		return dto.ExercisesDto{}, fmt.Errorf("failed to update workout: %w", err)
	}
	return exercise.ToDto(), nil
}
