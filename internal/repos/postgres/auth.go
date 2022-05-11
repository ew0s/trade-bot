package postgres

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/ew0s/trade-bot/internal/domain/entities"
	"github.com/ew0s/trade-bot/internal/repos/postgres/schema"
	"github.com/jmoiron/sqlx"
)

type auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *auth {
	return &auth{db: db}
}

func (r *auth) CreateUser(ctx context.Context, user entities.User) (string, error) {
	q := goqu.
		New("postgres", r.db).
		Insert(user).
		Into(schema.Users).
		Returning(goqu.C("uid")).
		Prepared(true).
		Executor()

	var uid string

	if _, err := q.ScanValContext(ctx, &uid); err != nil {
		return "", fmt.Errorf("executing query: %w", err)
	}

	return uid, nil
}

func (r *auth) GetUserByUID(ctx context.Context, uid string) (entities.User, bool, error) {
	q := goqu.New("postgres", r.db).
		From(goqu.T(schema.Users)).
		Select(entities.User{}).
		Where(goqu.C("uid").Eq(uid)).
		Prepared(true).
		Executor()

	var user entities.User

	found, err := q.ScanStructContext(ctx, &user)
	if err != nil {
		return entities.User{}, false, fmt.Errorf("scanning struct on query exec: %w", err)
	}

	return user, found, nil
}
