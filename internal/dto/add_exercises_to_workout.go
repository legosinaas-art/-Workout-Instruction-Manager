package dto

import "example.com/m/v2/internal/model"

type AddExercisesToWorkoutRequest struct {
	Exercises []model.ExerciseInWorkout `json:"exercises"`
}
