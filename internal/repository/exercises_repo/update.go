package exercises_repo

import (
	"context"

	"example.com/m/v2/internal/dao"
)

func (r *Repo) GetByID(ctx context.Context, id int) (dao.ExercisesDao, error) {
	const query = `
		SELECT id, name, notes
		FROM exercises
        WHERE id =$1`
	var exercises dao.ExercisesDao
	if err := r.db.GetContext(ctx, &exercises, query, id); err != nil {
		return dao.ExercisesDao{}, err
	}
	return exercises, nil
}

func (r *Repo) Update(ctx context.Context, m dao.ExercisesDao) error {
	const query = `
		UPDATE exercises
		SET name = $1, notes = $2 
        WHERE id =$3`
	_, err := r.db.ExecContext(ctx, query, m.Name, m.Notes, m.Id)
	if err != nil {
		return err
	}
	return nil
}
