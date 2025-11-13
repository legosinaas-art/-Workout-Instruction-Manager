package exercises_repo

import (
	"context"
	"database/sql"
	"errors"

	"example.com/m/v2/internal/dao"
	"example.com/m/v2/internal/repository/repo_error"
)

func (r *Repo) GetById(ctx context.Context, id int) (dao.ExercisesDao, error) {
	const query = `
		SELECT id, name, notes, created_at
		FROM exercises
		WHERE id = $1`
	var exercises dao.ExercisesDao
	if err := r.db.GetContext(ctx, &exercises, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dao.ExercisesDao{}, repo_error.ErrNotFound
		}
		return dao.ExercisesDao{}, err
	}
	return exercises, nil
}
