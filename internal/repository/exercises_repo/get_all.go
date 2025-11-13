package exercises_repo

import (
	"context"
	"fmt"
	"strings"

	"example.com/m/v2/internal/dao"
)

func (r *Repo) GetAll(ctx context.Context, limit, offset int, searchString string) ([]dao.ExercisesDao, int, error) {
	// Базовая часть запроса
	baseQuery := "FROM exercises"
	whereClause := "" // Условие WHERE изначально пустое

	// 1. Условное добавление фильтра
	// Этим мы обрабатываем "если фильтр пустой, просто не применять его"
	if searchString != "" {
		// Используем ILIKE для поиска без учета регистра (в PostgreSQL).
		// Оборачиваем строку поиска в %...% для поиска по вхождению.
		safeSearch := fmt.Sprintf("%%%s%%", strings.ReplaceAll(searchString, "%", "\\%"))

		// Создаем условие WHERE для поиска по name ИЛИ notes
		whereClause = fmt.Sprintf(" WHERE name ILIKE '%s' OR notes ILIKE '%s'", safeSearch, safeSearch)
	}

	// 2. Запрос для получения общего количества элементов (Total Count)
	// Добавляем условие WHERE к запросу COUNT(*), чтобы считать только отфильтрованные записи
	countQuery := fmt.Sprintf("SELECT COUNT(*) %s %s", baseQuery, whereClause)
	var total int

	// Выполняем запрос на подсчет
	if err := r.db.GetContext(ctx, &total, countQuery); err != nil {
		return nil, 0, fmt.Errorf("failed to get total count with filter: %w", err)
	}

	// 3. Запрос для получения данных на текущей странице (Paginated Data)
	// Сборка полного запроса с фильтром, сортировкой, LIMIT и OFFSET
	query := fmt.Sprintf(`
		SELECT id, name, notes, created_at
		%s %s 
		ORDER BY created_at DESC 
		LIMIT %d OFFSET %d`, baseQuery, whereClause, limit, offset)

	var exercises []dao.ExercisesDao
	// Выполняем запрос на выборку данных
	if err := r.db.SelectContext(ctx, &exercises, query); err != nil {
		return nil, 0, fmt.Errorf("failed to select exercises with filter: %w", err)
	}

	// 4. Возвращаем 3 значения: данные, общее количество, nil ошибки
	return exercises, total, nil
}
