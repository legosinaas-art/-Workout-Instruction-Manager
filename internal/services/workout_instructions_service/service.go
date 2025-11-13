package workout_instructions_service

import (
	"context"

	"example.com/m/v2/internal/dao"
	"example.com/m/v2/internal/framework"
	"example.com/m/v2/internal/model"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	repo          WorkoutInstructionsRepo
	exercisesRepo ExercisesRepo
	txRunner      framework.TxRunner
}

func NewService(repo WorkoutInstructionsRepo, exercisesRepo ExercisesRepo, txRunner framework.TxRunner) *Service {
	return &Service{repo: repo,
		exercisesRepo: exercisesRepo,
		txRunner:      txRunner}
}

type WorkoutInstructionsRepo interface {
	Create(ctx context.Context, m dao.WorkoutDao) (int, error)
	GetByID(ctx context.Context, id int) (dao.WorkoutDao, error)
	Update(ctx context.Context, m dao.WorkoutDao) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (dao.WorkoutDao, error)
	GetAll(ctx context.Context) ([]dao.WorkoutDao, error)
	GetByIdWithTx(ctx context.Context, tx *sqlx.Tx, id int) (dao.WorkoutDao, error)
	AddManyExercisesToInstructionWithTx(ctx context.Context, tx *sqlx.Tx, workoutId int, exercises []model.ExerciseInWorkout) error
	GetExercisesForInstruction(ctx context.Context, workoutId int) ([]model.ExerciseInWorkout, error)
}
type ExercisesRepo interface {
	GetById(ctx context.Context, id int) (dao.ExercisesDao, error)
	AreAllExercisesExistWithTx(ctx context.Context, tx *sqlx.Tx, exerciseIDs []int) error
}
