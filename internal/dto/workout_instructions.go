package dto

import (
	"time"

	"example.com/m/v2/internal/model"
)

// CreateWorkoutRequest - DTO для создания тренировки
type CreateWorkoutInstructionsRequest struct {
	Name  *string `json:"name"`
	Notes *string `json:"notes,omitempty"`
}

type CreateWorkoutInstructionsResponse struct {
	Id int `json:"id"`
}

type WorkoutInstructionsDto struct {
	Id        int                       `json:"id"`
	Name      *string                   `json:"name"`
	Notes     *string                   `json:"notes,omitempty"`
	CreatedAt time.Time                 `json:"created_at"`
	Exercises []model.ExerciseInWorkout `json:"exercises,omitempty"`
}
