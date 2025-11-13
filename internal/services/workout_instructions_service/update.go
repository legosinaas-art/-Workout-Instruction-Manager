package workout_instructions_service

import (
	"context"
	"fmt"

	"example.com/m/v2/internal/dto"
)

func (s *Service) Update(ctx context.Context, req dto.WorkoutInstructionsDto) (dto.WorkoutInstructionsDto, error) {
	workout, err := s.repo.GetByID(ctx, req.Id)
	if err != nil {
		return dto.WorkoutInstructionsDto{}, fmt.Errorf("not found id %w", err)
	}
	workout.Name = req.Name
	workout.Notes = req.Notes
	err = s.repo.Update(ctx, workout)
	if err != nil {
		return dto.WorkoutInstructionsDto{}, fmt.Errorf("failed to update workout: %w", err)
	}
	return workout.ToDto(), nil
}
