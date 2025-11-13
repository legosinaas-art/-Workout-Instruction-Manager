package workout_instructions_repo

import (
	"context"

	"example.com/m/v2/internal/dao"
)

func (r *Repo) GetByID(ctx context.Context, id int) (dao.WorkoutDao, error) {
	const query = `
		SELECT id, name, notes
		FROM workout_instructions 
        WHERE id =$1`
	var workout dao.WorkoutDao
	if err := r.db.GetContext(ctx, &workout, query, id); err != nil {
		return dao.WorkoutDao{}, err
	}
	return workout, nil
}

func (r *Repo) Update(ctx context.Context, m dao.WorkoutDao) error {
	const query = `
		UPDATE workout_instructions
		SET name = $1, notes = $2 
        WHERE id =$3`
	_, err := r.db.ExecContext(ctx, query, m.Name, m.Notes, m.Id)
	if err != nil {
		return err
	}
	return nil
}
