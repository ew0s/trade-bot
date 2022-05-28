package internal

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"

	"github.com/ew0s/trade-bot/internal/repos/postgres/models"
	"github.com/ew0s/trade-bot/internal/repos/postgres/schema"
	"github.com/ew0s/trade-bot/pkg/postgres"
)

func CreateUser(ctx context.Context, db postgres.QueryBuilder, user models.User) (string, error) {
	q := db.
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

func GetUserByUID(ctx context.Context, db postgres.QueryBuilder, uid string) (models.User, bool, error) {
	q := db.
		From(goqu.T(schema.Users)).
		Select(models.User{}).
		Where(
			goqu.C("uid").Eq(uid),
		).
		Prepared(true).
		Executor()

	var user models.User

	found, err := q.ScanStructContext(ctx, &user)
	if err != nil {
		return models.User{}, false, fmt.Errorf("scanning struct on query exec: %w", err)
	}

	return user, found, nil
}

func GetUserByUsername(ctx context.Context, db postgres.QueryBuilder, username string) (models.User, bool, error) {
	q := db.
		From(goqu.T(schema.Users)).
		Select(models.User{}).
		Where(
			goqu.C("username").Eq(username),
		).
		Prepared(true).
		Executor()

	var user models.User

	found, err := q.ScanStructContext(ctx, &user)
	if err != nil {
		return models.User{}, false, fmt.Errorf("scaning struct in query exec: %w", err)
	}

	return user, found, nil
}
