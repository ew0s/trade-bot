package postgres

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/ew0s/trade-bot/pkg/constant"
	"github.com/jmoiron/sqlx"

	"github.com/ew0s/trade-bot/internal/domain/entities"
	"github.com/ew0s/trade-bot/internal/repos/postgres/schema"
)

type auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *auth {
	return &auth{db: db}
}

func (r *auth) CreateUser(ctx context.Context, user entities.User) (string, error) {
	_, found, err := GetUserByUsername(ctx, r.db, user.Username)
	if err != nil {
		return "", fmt.Errorf("getting user by username: %w", err)
	}

	if found {
		return "", fmt.Errorf("getting user by username: already exists: %w", constant.ErrBadRequest)
	}

	q := goqu.
		New("postgres", r.db).
		Insert(goqu.T(schema.Users)).
		Rows(user).
		Prepared(true).
		Returning(goqu.C("uid")).
		Executor()

	var uid string

	if _, err := q.ScanValContext(ctx, &uid); err != nil {
		return "", fmt.Errorf("executing query: %w", err)
	}

	return uid, nil
}

func (r *auth) GetUserByUID(ctx context.Context, uid string) (entities.User, bool, error) {
	return GetUserByUID(ctx, r.db, uid)
}

func GetUserByUID(ctx context.Context, db *sqlx.DB, uid string) (entities.User, bool, error) {
	q := goqu.New("postgres", db).
		From(goqu.T(schema.Users)).
		Select(entities.User{}).
		Where(
			goqu.C("uid").Eq(uid),
		).
		Prepared(true).
		Executor()

	var user entities.User

	found, err := q.ScanStructContext(ctx, &user)
	if err != nil {
		return entities.User{}, false, fmt.Errorf("scanning struct on query exec: %w", err)
	}

	return user, found, nil
}

func GetUserByUsername(ctx context.Context, db *sqlx.DB, username string) (entities.User, bool, error) {
	q := goqu.
		New("postgres", db).
		From(goqu.T(schema.Users)).
		Select(entities.User{}).
		Where(
			goqu.C("username").Eq(username),
		).
		Prepared(true).
		Executor()

	var user entities.User

	found, err := q.ScanStructContext(ctx, &user)
	if err != nil {
		return entities.User{}, false, fmt.Errorf("scaning struct in query exec: %w", err)
	}

	return user, found, nil
}
