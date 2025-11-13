package workout_instructions_service

import (
	"context"
	"errors"
	"fmt"

	"example.com/m/v2/internal/model"
	"example.com/m/v2/internal/repository/repo_error"
	"github.com/jmoiron/sqlx"
)

func (s *Service) AddExercises(
	ctx context.Context,
	workoutInstructionsId int,
	exercises []model.ExerciseInWorkout) error {

	err := s.txRunner.RunInTx(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		_, err := s.repo.GetByIdWithTx(ctx, tx, workoutInstructionsId)
		if err != nil {
			if errors.Is(err, repo_error.ErrNotFound) {
				return repo_error.ErrNotFound // <-- Правильно "пробрасываем"
			}
			return fmt.Errorf("s.repo.GetByIdWithTx: %w", err)
		}
		var exerciseIDs []int
		for _, ids := range exercises {
			exerciseIDs = append(exerciseIDs, ids.Id)
		}
		err = s.exercisesRepo.AreAllExercisesExistWithTx(ctx, tx, exerciseIDs)
		if err != nil {
			return fmt.Errorf("s.exercisesRepo.AreAllExercisesExistWithTx:%w", err)
		}

		err = s.repo.AddManyExercisesToInstructionWithTx(ctx, tx, workoutInstructionsId, exercises)
		if err != nil {
			return fmt.Errorf("s.repo.AddManyExercisesToInstruction:%w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("txRunner.RunInTx: %w", err)
	}
	return nil
}
