package postgres

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	UsersTable = "users"
	WorkTable  = "workout_instructions_repo"
	StatsTable = "stats"
	ExeTable   = "exercises"
	MealsTable = "meals"
)

func NewPostgresDB(dsn string) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "pgx", dsn)

	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	return db, nil
}
