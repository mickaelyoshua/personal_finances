package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Agent interface {
	Querier
}

type SQLAgent struct {
	Conn    *pgx.Conn
	*Queries
}

func NewAgent(ctx context.Context, databaseURL string) (Agent, error) {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	return &SQLAgent{
		Conn:    conn,
		Queries: New(conn),
	}, nil
}