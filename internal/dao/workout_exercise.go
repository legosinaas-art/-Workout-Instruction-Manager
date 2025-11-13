package dao

import (
	"example.com/m/v2/internal/framework/jsonb"
	"example.com/m/v2/internal/model"
)

type WorkoutExercise struct {
	workoutId       int                                `db:"workout_id"`
	exerciseId      int                                `db:"exercise_id"`
	orderNum        int                                `db:"order_num"`
	exerciseDetails jsonb.JSONB[model.ExerciseDetails] `db:"details"`
}
