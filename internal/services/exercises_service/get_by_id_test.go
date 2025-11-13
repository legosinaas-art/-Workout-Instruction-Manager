// Файл: internal/services/exercises_service/get_by_id_test.go
package exercises_service

import (
	"context"
	"errors"  // Нам понадобится для 'другой' (фейковой) ошибки
	"testing" // Главный пакет для тестов
	"time"    // Для создания 'фейковых' данных

	"example.com/m/v2/internal/dao"
	"example.com/m/v2/internal/repository/repo_error"
	// `assert` - это "помощник" (из `testify`), который делает проверки (типа "я ожидаю, что А = Б")
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock" // Нам нужен 'mock' для `mock.Anything`
)

// TestService_GetById - это НАШ ТЕСТ.
// Названия тестов в Go всегда начинаются с `Test...`
func TestService_GetById(t *testing.T) {

	// --- 1. АРАНЖИРОВКА (Arrange) ---
	// Мы "строим сцену" и "раздаем роли" перед началом "пьесы".

	// Создаем нашего "Актера" (Мок-репозиторий, который мы написали в `repo_mock_test.go`)
	mockRepo := new(MockExercisesRepo)

	// Создаем "Менеджера" (наш Сервис), но "подсовываем" ему
	// нашего "Актера" (`mockRepo`) вместо настоящего репозитория.
	service := NewService(mockRepo)

	// Создаем "фейковые" данные, которые мы *ожидаем*.
	// Это то, что "Актер" (Мок) "вернет" нам якобы из "базы".
	mockName := "Приседания"
	mockNotes := "Глубокий сед"
	mockExerciseDAO := dao.ExercisesDao{
		Id:        1,
		Name:      &mockName,  // Используем & (указатель), т.к. в DAO у нас *string
		Notes:     &mockNotes, // Используем & (указатель)
		CreatedAt: time.Now(),
	}

	// --- 2. ДЕЙСТВИЕ (Act) ---
	// Мы "разыгрываем" наши "сценарии" (тест-кейсы).
	// В Go для этого очень удобно использовать `t.Run(...)`

	// --- Сценарий 1: Успех ---
	t.Run("Успех - Упражнение найдено", func(t *testing.T) {

		// --- Сценарий для "Актера" ---
		// Мы "говорим" нашему "Актеру" (Моку):
		// "Когда тебя вызовут (`.On`) с методом `GetById`...
		//  с любым контекстом (`mock.Anything`) и ID=1...
		mockRepo.On("GetById", mock.Anything, 1).
			//  ...ты ДОЛЖЕН вернуть (`.Return`) `mockExerciseDAO` и `nil` (нет ошибки)."
			Return(mockExerciseDAO, nil).
			Once() // `.Once()` - хорошая привычка, "ожидаем, что тебя вызовут 1 раз"

		// --- Запускаем наш метод ---
		// Вызываем "Менеджера" (Сервис) и просим найти ID=1.
		// Это и есть само **Действие**, которое мы тестируем.
		dto, err := service.GetById(context.Background(), 1)

		// --- 3. ПРОВЕРКА (Assert) ---
		// Мы "проверяем" результат.

		// "Я утверждаю (assert), что ошибки (err) НЕТ (Nil)."
		assert.Nil(t, err)
		// "Я утверждаю, что DTO (результат) НЕ пустой (NotNil)."
		assert.NotNil(t, dto)
		// "Я утверждаю, что ID в DTO равен 1."
		assert.Equal(t, 1, dto.Id)
		// "Я утверждаю, что Имя в DTO - 'Приседания'."
		assert.Equal(t, "Приседания", *dto.Name) // *dto.Name, т.к. `Name` - это *string
	})

	// --- Сценарий 2: Не найдено ---
	t.Run("Ошибка - Упражнение не найдено (ErrNotFound)", func(t *testing.T) {

		// --- Сценарий для "Актера" ---
		// "Когда тебя вызовут (`.On`) с `GetById`, любым контекстом и ID=99...
		mockRepo.On("GetById", mock.Anything, 99).
			//  ...ты ДОЛЖЕН вернуть (`.Return`) пустой `dao.ExercisesDao{}`...
			//  ...и ошибку `repo_error.ErrNotFound`."
			Return(dao.ExercisesDao{}, repo_error.ErrNotFound).
			Once()

		// --- Запускаем наш метод ---
		dto, err := service.GetById(context.Background(), 99)

		// --- 3. ПРОВЕРКА (Assert) ---
		// "Я утверждаю, что ошибка (err) ЕСТЬ (NotNil)."
		assert.NotNil(t, err)
		// "Я утверждаю, что DTO (результат) - пустой (Id == 0)."
		assert.Equal(t, 0, dto.Id)

		// "Я утверждаю (проверяю), что ошибка (err) - это *именно* `repo_error.ErrNotFound`."
		// Это тот самый "проброс", который мы с тобой чинили в `get_by_id.go`!
		assert.ErrorIs(t, err, repo_error.ErrNotFound)
	})

	// --- Сценарий 3: Ошибка базы данных ---
	t.Run("Ошибка - Другая ошибка (БД отвалилась)", func(t *testing.T) {

		// Создаем "фейковую" ошибку
		dbError := errors.New("БД отвалилась")

		// --- Сценарий для "Актера" ---
		// "Когда тебя вызовут (`.On`) с `GetById`, любым контекстом и ID=500...
		mockRepo.On("GetById", mock.Anything, 500).
			//  ...ты ДОЛЖЕН вернуть (`.Return`) пустой `dao.ExercisesDao{}`...
			//  ...и нашу 'фейковую' ошибку `dbError`."
			Return(dao.ExercisesDao{}, dbError).
			Once()

		// --- Запускаем наш метод ---
		dto, err := service.GetById(context.Background(), 500)

		// --- 3. ПРОВЕРКА (Assert) ---
		// "Я утверждаю, что ошибка (err) ЕСТЬ (NotNil)."
		assert.NotNil(t, err)
		// "Я утверждаю, что DTO - пустой."
		assert.Equal(t, 0, dto.Id)
		// "Я утверждаю, что ошибка (err) - это *НЕ* `repo_error.ErrNotFound`."
		assert.NotErrorIs(t, err, repo_error.ErrNotFound)

		// "Я утверждаю, что текст ошибки (err.Error()) *содержит* (Contains) 'БД отвалилась'."
		// Этим мы проверяем, что наш сервис "завернул" ошибку, как и должен.
		assert.Contains(t, err.Error(), "БД отвалилась")
	})
}
