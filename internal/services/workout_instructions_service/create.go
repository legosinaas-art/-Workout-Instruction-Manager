package workout_instructions_service

import (
	"context"
	"fmt"

	"example.com/m/v2/internal/dao"
	"example.com/m/v2/internal/dto"
)

func (s *Service) Create(ctx context.Context, req dto.CreateWorkoutInstructionsRequest) (dto.CreateWorkoutInstructionsResponse, error) {
	id, err := s.repo.Create(ctx, dao.WorkoutDao{
		Name:  req.Name,
		Notes: req.Notes,
	})
	if err != nil {
		return dto.CreateWorkoutInstructionsResponse{}, fmt.Errorf("s.repo.Create: %w", err)
	}
	return dto.CreateWorkoutInstructionsResponse{
		Id: id,
	}, nil
}
