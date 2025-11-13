package dao

import (
	"time"

	"example.com/m/v2/internal/dto"
)

// WorkoutDao - основная модель тренировки для базы данных
type WorkoutDao struct {
	Id        int       `db:"id"`
	Name      *string   `db:"name"`
	Notes     *string   `db:"notes"`
	CreatedAt time.Time `db:"created_at"`
}

func (workout WorkoutDao) ToDto() dto.WorkoutInstructionsDto {
	return dto.WorkoutInstructionsDto{
		Id:        workout.Id,
		Name:      workout.Name,
		Notes:     workout.Notes,
		CreatedAt: workout.CreatedAt,
	}
}
