package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TxDoer interface {
	DoInTransaction(ctx context.Context, opts *sql.TxOptions, fn func(tx *sql.Tx) error) error
}

type txDoer struct {
	db *sqlx.DB
}

func NewTxDoer(db *sqlx.DB) *txDoer {
	return &txDoer{db}
}

func (t *txDoer) DoInTransaction(ctx context.Context, opts *sql.TxOptions, fn func(tx *sql.Tx) error) error {
	tx, err := t.db.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("creating transaction: %w", err)
	}

	if txErr := fn(tx); txErr != nil {
		if err = tx.Rollback(); err != nil {
			return fmt.Errorf("rolling transaction back: %w", err)
		}

		return fmt.Errorf("executing in transaction: %w", txErr)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}
