package exercises_repo

import (
	"context"

	"example.com/m/v2/internal/dao"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

type Repository interface {
	Create(ctx context.Context, m dao.ExercisesDao) (int, error)
	GetByID(ctx context.Context, id int) (dao.ExercisesDao, error)
	Update(ctx context.Context, m dao.ExercisesDao) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (dao.ExercisesDao, error)
	GetAll(ctx context.Context, limit, offset int, searchString string) ([]dao.ExercisesDao, int, error)
	AreAllExercisesExistWithTx(ctx context.Context, tx sqlx.Tx, exerciseIDs []int) error
}
