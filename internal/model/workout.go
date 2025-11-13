package model

import "time"

// Workout - основная модель тренировки
type Workout struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	Date        time.Time `json:"date" db:"date"`
	Title       string    `json:"title" db:"title"`
	Notes       *string   `json:"notes,omitempty" db:"notes"` // указатель для nullable поля
	DurationMin *int      `json:"duration_min,omitempty" db:"duration_min"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
