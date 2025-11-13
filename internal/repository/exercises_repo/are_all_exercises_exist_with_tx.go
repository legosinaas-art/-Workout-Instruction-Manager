package exercises_repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (r *Repo) AreAllExercisesExistWithTx(
	ctx context.Context,
	tx *sqlx.Tx,
	exerciseIDs []int) error {
	if len(exerciseIDs) == 0 {
		return nil
	}
	uniqueIDs := make(map[int]struct{})
	for _, id := range exerciseIDs {
		uniqueIDs[id] = struct{}{}
	}
	expectedCount := len(uniqueIDs)
	query, args, err := sqlx.In("SELECT COUNT(DISTINCT id) FROM exercises WHERE id IN (?)", exerciseIDs)
	if err != nil {
		return fmt.Errorf("не удалось собрать SQL-запрос для IN: %w", err)
	}
	// 4. `sqlx` по умолчанию использует `?`, а Postgres ждет `$1`.
	// `tx.Rebind` "чинит" запрос.
	query = tx.Rebind(query)
	var countInDB int
	if err = tx.GetContext(ctx, &countInDB, query, args...); err != nil {
		return fmt.Errorf("ошибка при выполнении COUNT: %w", err)
	}

	if countInDB != expectedCount {
		return fmt.Errorf("одно или несколько упражнений не найдены (ожидалось: %d, найдено: %d)", expectedCount, countInDB)
	}
	return nil
}
