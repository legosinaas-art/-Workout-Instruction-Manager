package workout_instructions_repo

import (
	"context"
	"database/sql"
	"errors"

	"example.com/m/v2/internal/dao"
	"example.com/m/v2/internal/repository/repo_error"
	"github.com/jmoiron/sqlx"
)

func (r *Repo) GetByIdWithTx(ctx context.Context, tx *sqlx.Tx, id int) (dao.WorkoutDao, error) {
	const query = `
		SELECT id, name, notes, created_at
		FROM workout_instructions
		WHERE id = $1`
	var workout dao.WorkoutDao
	if err := tx.GetContext(ctx, &workout, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dao.WorkoutDao{}, repo_error.ErrNotFound
		}
		return dao.WorkoutDao{}, err
	}
	return workout, nil
}
