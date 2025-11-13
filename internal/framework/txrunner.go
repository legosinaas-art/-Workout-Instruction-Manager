package framework

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TxRunner struct {
	db *sqlx.DB
}

func NewTxRunner(db *sqlx.DB) *TxRunner {
	return &TxRunner{db: db}
}
func (t *TxRunner) RunInTx(ctx context.Context, run func(ctx context.Context, tx *sqlx.Tx) error) error {
	tx, err := t.db.Beginx()
	if err != nil {
		return fmt.Errorf("error to run transaction: %w", err)
	}
	defer func() {
		switch p := recover(); {
		case p != nil:
			err = tx.Rollback()
			if err != nil {
				err = fmt.Errorf("panica unable rollback tx:%w", err)
			}
			err = fmt.Errorf("recovered transaction from panic: %v", p)
		case err != nil:
			err = tx.Rollback()
			if err != nil {
				err = fmt.Errorf("unable rollback tx:%w", err)
			}
		default:
			if err = tx.Commit(); err != nil {
				err = fmt.Errorf("cannot commit transaction: %w", err)
			}
		}
	}()
	err = run(ctx, tx)
	return err
}
