package workout_instructions_service

import (
	"context"
	"fmt"

	"example.com/m/v2/internal/dto"
)

func (s *Service) GetAll(ctx context.Context) ([]dto.WorkoutInstructionsDto, error) {
	workouts, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to show workout: %w", err)
	}
	result := make([]dto.WorkoutInstructionsDto, 0, len(workouts))
	for _, workout := range workouts {
		result = append(result, workout.ToDto())
	}
	return result, nil
}
