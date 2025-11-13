package workout_instructions_service

import (
	"context"
	"errors"
	"fmt"

	"example.com/m/v2/internal/dto"
	"example.com/m/v2/internal/repository/repo_error"
)

func (s *Service) GetById(ctx context.Context, id int) (dto.WorkoutInstructionsDto, error) {

	workout, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, repo_error.ErrNotFound) {
			return dto.WorkoutInstructionsDto{}, repo_error.ErrNotFound
		}
		return dto.WorkoutInstructionsDto{}, fmt.Errorf("s.repo.GetById: %w", err)
	}

	//---------------------------------------------------------------
	exercises, err := s.repo.GetExercisesForInstruction(ctx, id)
	if err != nil {
		return dto.WorkoutInstructionsDto{}, fmt.Errorf("s.repo.GetExercisesForInstruction: %w", err)
	}
	resultDto := dto.WorkoutInstructionsDto{
		Id:        workout.Id,
		Name:      workout.Name,
		Notes:     workout.Notes,
		CreatedAt: workout.CreatedAt,
		Exercises: exercises,
	}
	return resultDto, nil
}
