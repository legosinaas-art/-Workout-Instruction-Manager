package dao

import (
	"time"

	"example.com/m/v2/internal/dto"
)

type ExercisesDao struct {
	Id        int       `db:"id"`
	Name      *string   `db:"name"`
	Notes     *string   `db:"notes"`
	CreatedAt time.Time `db:"created_at"`
}

func (exercises ExercisesDao) ToDto() dto.ExercisesDto {
	return dto.ExercisesDto{
		Id:        exercises.Id,
		Name:      exercises.Name,
		Notes:     exercises.Notes,
		CreatedAt: exercises.CreatedAt,
	}
}
