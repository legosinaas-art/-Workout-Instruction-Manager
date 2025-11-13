package exercises_service

import (
	"context"

	"example.com/m/v2/internal/dao"
	// `stretchr/testify` - это стандартная библиотека для тестов,
	// она нам понадобится, но Go ее сам импортирует, когда мы начнем писать тест.
	"github.com/stretchr/testify/mock"
)

// MockExercisesRepo - это наша "заглушка" ("Актер")
// Он "притворяется" репозиторием.
type MockExercisesRepo struct {
	mock.Mock // Он "наследует" поведение из `testify/mock`
}

// --- Теперь мы реализуем *все* методы из интерфейса ExercisesRepo ---
// (Интерфейс ExercisesRepo описан у тебя в `service.go`)

// GetById - это тот метод, который нам нужен для теста
func (m *MockExercisesRepo) GetById(ctx context.Context, id int) (dao.ExercisesDao, error) {
	// Эта магия `testify/mock` означает:
	// 1. "Запиши, что меня вызвали с этими аргументами (ctx, id)"
	args := m.Called(ctx, id)

	// 2. "Верни то, что мы настроили в 'сценарии'":
	//    - Первое значение (`args.Get(0)`) - это `dao.ExercisesDao`
	//    - Второе значение (`args.Error(1)`) - это `error`
	return args.Get(0).(dao.ExercisesDao), args.Error(1)
}

// --- Это "пустышки" для остальных методов интерфейса. Они нам не нужны для этого теста, ---
// --- но они *обязаны* быть, чтобы наш Мок "удовлетворял" интерфейсу. ---

func (m *MockExercisesRepo) GetByID(ctx context.Context, id int) (dao.ExercisesDao, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(dao.ExercisesDao), args.Error(1)
}

func (m *MockExercisesRepo) Create(ctx context.Context, model dao.ExercisesDao) (int, error) {
	args := m.Called(ctx, model)
	return args.Int(0), args.Error(1)
}

func (m *MockExercisesRepo) Update(ctx context.Context, model dao.ExercisesDao) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockExercisesRepo) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockExercisesRepo) GetAll(ctx context.Context, limit, offset int, searchString string) ([]dao.ExercisesDao, int, error) {
	args := m.Called(ctx, limit, offset, searchString)
	return args.Get(0).([]dao.ExercisesDao), args.Int(1), args.Error(2)
}
