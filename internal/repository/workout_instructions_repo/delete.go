package workout_instructions_repo

import (
	"context"
)

func (r *Repo) Delete(ctx context.Context, id int) error {
	const query = `
		DELETE FROM workout_instructions
		WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
