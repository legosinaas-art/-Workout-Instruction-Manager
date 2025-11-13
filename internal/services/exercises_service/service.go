package exercises_service

import (
	"context"

	"example.com/m/v2/internal/dao"
)

type Service struct {
	repo ExercisesRepo
}

func NewService(repo ExercisesRepo) *Service {
	return &Service{repo: repo}
}

type ExercisesRepo interface {
	Create(ctx context.Context, m dao.ExercisesDao) (int, error)
	GetByID(ctx context.Context, id int) (dao.ExercisesDao, error)
	Update(ctx context.Context, m dao.ExercisesDao) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (dao.ExercisesDao, error)
	GetAll(ctx context.Context, limit, offset int, searchString string) ([]dao.ExercisesDao, int, error)
}
