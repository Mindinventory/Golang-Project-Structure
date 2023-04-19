package repository

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"

	"golang-program-structure/common/database"
)

type GetUserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

func GetUserById(ctx context.Context, userId *uuid.UUID) (GetUserResponse, error) {
	var user GetUserResponse

	p, err := database.GetConnectionPool()
	if err != nil {
		return user, err
	}

	c, err := p.Acquire(ctx)
	if err != nil {
		return user, err
	}
	defer c.Release()

	tx, err := c.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return user, err
	}

	rows, err := tx.Query(ctx, `
		SELECT *
		FROM users
		WHERE user_id=$1
	`, userId)
	if err != nil {
		return user, err
	}

	err = rows.Scan(&user)
	if err != nil {
		tx.Rollback(ctx)
		return user, err
	}

	return user, nil
}
