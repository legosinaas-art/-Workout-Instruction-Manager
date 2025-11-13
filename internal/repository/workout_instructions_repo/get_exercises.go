package workout_instructions_repo

import (
	"context"
	"fmt"

	"example.com/m/v2/internal/framework/jsonb"
	"example.com/m/v2/internal/model"
)

func (r *Repo) GetExercisesForInstruction(ctx context.Context, workoutId int) ([]model.ExerciseInWorkout, error) {
	type workoutExerciseRow struct {
		ExerciseId      int                                `db:"exercise_id"`
		OrderNum        int                                `db:"order_num"`
		ExerciseDetails jsonb.JSONB[model.ExerciseDetails] `db:"details"`
	}

	const query = `
		SELECT exercise_id, order_num, details 
		FROM workout_instructions_exercises
		WHERE workout_instruction_id = $1
		ORDER BY order_num ASC`

	var rows []workoutExerciseRow

	err := r.db.SelectContext(ctx, &rows, query, workoutId)

	if err != nil {
		return nil, fmt.Errorf("r.db.SelectContext %d: %w", workoutId, err)
	}

	result := make([]model.ExerciseInWorkout, 0, len(rows))

	for _, row := range rows {
		details, _ := row.ExerciseDetails.GetFull()
		cleanExercise := model.ExerciseInWorkout{
			Id:              row.ExerciseId,
			ExerciseDetails: details,
			OrderNum:        row.OrderNum,
		}
		result = append(result, cleanExercise)
	}

	return result, nil

}
