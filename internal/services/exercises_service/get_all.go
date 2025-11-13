package exercises_service

import (
	"context"
	"fmt"

	"example.com/m/v2/internal/dto"
)

func (s *Service) GetAll(ctx context.Context, req dto.GetExercisesPaginationRequest) (dto.PaginatedResponse[dto.ExercisesDto], error) {
	page := req.Pagination.Page
	perPage := req.Pagination.PerPage
	searchString := req.Filter.Search

	offset := (page - 1) * perPage
	limit := perPage

	exercises, total, err := s.repo.GetAll(ctx, limit, offset, searchString)
	if err != nil {
		return dto.PaginatedResponse[dto.ExercisesDto]{}, fmt.Errorf("failed to show workout: %w", err)
	}
	result := make([]dto.ExercisesDto, 0, len(exercises))
	for _, exercise := range exercises {
		result = append(result, exercise.ToDto())
	}
	paginatedResponse := dto.NewPaginatedResponse(
		result,
		page,
		perPage,
		total,
	)
	return paginatedResponse, nil
}
