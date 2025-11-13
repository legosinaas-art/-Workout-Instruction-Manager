package exercises_service

import (
	"context"
	"fmt"

	"example.com/m/v2/internal/dao"
	"example.com/m/v2/internal/dto"
)

func (s *Service) Create(ctx context.Context, req dto.CreateExercisesRequest) (dto.CreateExercisesResponse, error) {
	id, err := s.repo.Create(ctx, dao.ExercisesDao{
		Name:  req.Name,
		Notes: req.Notes,
	})
	if err != nil {
		return dto.CreateExercisesResponse{}, fmt.Errorf("s.repo.Create: %w", err)
	}
	return dto.CreateExercisesResponse{
		Id: id,
	}, nil
}
