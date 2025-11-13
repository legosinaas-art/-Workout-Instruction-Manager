package exercises_repo

import (
	"context"
)

func (r *Repo) Delete(ctx context.Context, id int) error {
	const query = `
		DELETE FROM exercises
		WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
