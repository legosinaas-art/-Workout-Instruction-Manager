// Файл: internal/framework/jsonb/jsonb.go
package jsonb

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/samber/lo"
)

// JSONB универсальный тип для работы с jsonb в PostgreSQL
type JSONB[T any] struct {
	data  T
	valid bool // Для обработки NULL значений
}

func NewFromPtr[T any](value *T) JSONB[T] {
	if value != nil {
		return JSONB[T]{
			data:  *value,
			valid: true,
		}
	}

	return JSONB[T]{}
}

func New[T any](value T) JSONB[T] {
	return JSONB[T]{
		data:  value,
		valid: true,
	}
}

func (j *JSONB[T]) GetPtr() *T {
	if j.valid {
		return &j.data
	}

	return nil
}

func (j *JSONB[T]) GetFull() (T, bool) {
	return j.data, j.valid
}

func (j *JSONB[T]) Get() T {
	if j.valid {
		return j.data
	}

	return lo.Empty[T]()
}

// Scan реализует интерфейс sql.Scanner
func (j *JSONB[T]) Scan(value interface{}) error {
	if value == nil {
		j.valid = false
		return nil
	}

	var bytes []byte // <-- Мы будем работать с байтами

	// --- ВОТ ГЛАВНЫЙ ФИКС ---
	// "Умная" проверка: что бы ни пришло (string или []byte),
	// мы превращаем это в []byte.
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		// А вот тут мы починили твою опечатку в ошибке
		return fmt.Errorf("%s.Scan: unsupported type %T for Scan", j.typeName(), value)
	}
	// -----------------------

	if err := json.Unmarshal(bytes, &j.data); err != nil {
		j.valid = false
		return fmt.Errorf("%s.Scan: %w", j.typeName(), err)
	}
	j.valid = true

	return nil
}

// Value реализует интерфейс driver.Valuer
func (j JSONB[T]) Value() (driver.Value, error) {
	if !j.valid {
		return nil, nil
	}
	extData, err := json.Marshal(j.data)
	if err != nil {
		return nil, fmt.Errorf("%s.Value: %w", j.typeName(), err)
	}

	// Отдаем в базу тоже как []byte (это стандарт)
	return extData, nil
}

func (j JSONB[T]) typeName() string {
	return reflect.TypeOf(j.data).String()
}
