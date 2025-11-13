// Файл: internal/repository/workout_instructions_repo/add_exercises.go
package workout_instructions_repo

import (
	"context"
	"fmt"
	"strings"

	"example.com/m/v2/internal/framework/jsonb"
	"example.com/m/v2/internal/model"
	"github.com/jmoiron/sqlx"
)

// AddManyExercisesToInstruction добавляет сразу несколько упражнений в тренировку одним запросом.
func (r *Repo) AddManyExercisesToInstructionWithTx(
	ctx context.Context,
	tx *sqlx.Tx,
	workoutId int,
	exercises []model.ExerciseInWorkout,
) error {
	// 1. Проверяем, есть ли вообще что добавлять.
	// Если список упражнений пустой, то и делать ничего не нужно.
	if len(exercises) == 0 {
		return nil
	}

	// 2. Подготовка к "сборке" запроса.
	// Нам нужно превратить список упражнений в две вещи:
	//    - queryPlaceholders: Строку вида "($1, $2, $3, $4), ($5, $6, $7, $8), ..."
	//    - queryArgs: Плоский список всех значений [workoutId, ex1.Id, 1, ex1.details, workoutId, ex2.Id, 2, ex2.details, ...]

	queryPlaceholders := make([]string, 0, len(exercises)) // Сюда будем складывать "(?, ?, ?, ?)"
	queryArgs := make([]interface{}, 0, len(exercises)*4)  // А сюда - все значения для запроса

	// 3. Проходим по всем упражнениям в цикле, чтобы собрать части.
	for i, exercise := range exercises {
		// Превращаем детали (вес, повторения) в JSON-текст
		// Добавляем в `queryPlaceholders` шаблон для одной строки.
		// Вместо $1, $2... мы используем ?, sqlx потом сам заменит их на нужный диалект ($1 для Postgres).
		// i*4+1 - это магия, чтобы правильно рассчитать номера плейсхолдеров.
		// Для первого упражнения будет ($1, $2, $3, $4), для второго ($5, $6, $7, $8) и т.д.
		placeholder := fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4)
		queryPlaceholders = append(queryPlaceholders, placeholder)

		// Добавляем в `queryArgs` все значения для этой строки в правильном порядке.
		queryArgs = append(queryArgs, workoutId)                           // $1, $5, $9...
		queryArgs = append(queryArgs, exercise.Id)                         // $2, $6, $10...
		queryArgs = append(queryArgs, exercise.OrderNum)                   // $3, $7, $11... (Порядковый номер)
		queryArgs = append(queryArgs, jsonb.New(exercise.ExerciseDetails)) // $4, $8, $12...
	}

	// 4. Собираем финальный SQL-запрос.
	query := fmt.Sprintf(`
		INSERT INTO workout_instructions_exercises (
			workout_instruction_id, 
			exercise_id, 
			order_num, 
			details
		) VALUES %s`, strings.Join(queryPlaceholders, ", ")) // Соединяем все "(?, ?, ?, ?)" через запятую

	// 5. Выполняем один большой запрос.
	_, err := tx.ExecContext(ctx, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("ошибка при массовом добавлении упражнений в тренировку: %w", err)
	}

	// 6. Если все прошло хорошо, возвращаем "нет ошибки".
	return nil
}
