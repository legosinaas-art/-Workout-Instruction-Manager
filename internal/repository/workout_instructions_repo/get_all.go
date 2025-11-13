package workout_instructions_repo

import (
	"context"

	"example.com/m/v2/internal/dao"
)

func (r *Repo) GetAll(ctx context.Context) ([]dao.WorkoutDao, error) {
	const query = `
		SELECT id, name, notes, created_at
		FROM workout_instructions`
	var workout []dao.WorkoutDao
	if err := r.db.SelectContext(ctx, &workout, query); err != nil {
		return nil, err
	}
	return workout, nil
}
