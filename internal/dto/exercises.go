package dto

import (
	"time"
)

type GetExercisesPaginationRequest struct {
	Pagination Pagination
	Filter     Filter
}
type Pagination struct {
	Page    int
	PerPage int
}
type Filter struct {
	Search string
}

func (p *Pagination) SetDefaults() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 {
		p.PerPage = 10
	}
}

type CreateExercisesRequest struct {
	Name  *string `json:"name"`
	Notes *string `json:"notes,omitempty"`
}

type CreateExercisesResponse struct {
	Id int `json:"id"`
}

type ExercisesDto struct {
	Id        int       `json:"id"`
	Name      *string   `json:"name"`
	Notes     *string   `json:"notes,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
