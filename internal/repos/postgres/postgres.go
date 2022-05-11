package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ew0s/trade-bot/internal/configer/appcofig"
)

func NewPostgresDB(cfg appcofig.PostgresConfiguration) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", postgresDatasource(cfg))
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
