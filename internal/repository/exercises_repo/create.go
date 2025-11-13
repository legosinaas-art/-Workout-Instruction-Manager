package exercises_repo

import (
	"context"

	"example.com/m/v2/internal/dao"
)

func (r *Repo) Create(ctx context.Context, m dao.ExercisesDao) (int, error) {
	var id int
	const query = `
		INSERT INTO exercises (name, notes) 
		VALUES ($1, $2) RETURNING id`

	if err := r.db.QueryRowxContext(ctx, query, m.Name, m.Notes).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
