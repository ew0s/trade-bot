package postgres

import (
	"context"
	"fmt"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // set postgres dialect for goqu builder
	_ "github.com/jackc/pgx/stdlib"                     // driver for sql
	"github.com/jmoiron/sqlx"

	"github.com/ew0s/trade-bot/internal/configer/appcofig"
)

func NewPostgresDB(ctx context.Context, cfg appcofig.PostgresConfiguration) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "pgx", postgresDatasource(cfg))
	if err != nil {
		return nil, fmt.Errorf("connecting to postgres db: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging postgres db: %w", err)
	}

	return db, nil
}

func postgresDatasource(cfg appcofig.PostgresConfiguration) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Credentials.Username,
		cfg.Credentials.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
}
