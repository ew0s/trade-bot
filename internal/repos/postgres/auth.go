package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/ew0s/trade-bot/internal/domain/entities"
	"github.com/ew0s/trade-bot/internal/repos/postgres/internal"
	"github.com/ew0s/trade-bot/internal/repos/postgres/mapper"
	"github.com/ew0s/trade-bot/pkg/constant"
	"github.com/ew0s/trade-bot/pkg/postgres"
)

type auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *auth {
	return &auth{db: db}
}

func (r *auth) CreateUser(ctx context.Context, user entities.User) (string, error) {
	opts := &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	}

	var uid string

	if err := postgres.NewTxDoer(r.db).DoInTransaction(ctx, opts, func(tx *sql.Tx) error {
		db := postgres.NewTxBuilder(tx)

		_, found, err := internal.GetUserByUsername(ctx, db, user.Username)
		if err != nil {
			return fmt.Errorf("getting user by username: %w", err)
		}

		if found {
			return fmt.Errorf("getting user by username: already exists: %w", constant.ErrBadRequest)
		}

		uid, err = internal.CreateUser(ctx, db, mapper.MakeModelUser(user))
		if err != nil {
			return fmt.Errorf("creating user: %w", err)
		}

		return nil
	}); err != nil {
		return "", fmt.Errorf("doing in transaction: %w", err)
	}

	return uid, nil
}

func (r *auth) GetUserByUID(ctx context.Context, uid string) (entities.User, bool, error) {
	user, found, err := internal.GetUserByUID(ctx, postgres.NewBuilder(r.db.DB), uid)
	if err != nil {
		return entities.User{}, false, fmt.Errorf("getting user by uid: %w", err)
	}

	return mapper.MakeEntityUser(user), found, nil
}
