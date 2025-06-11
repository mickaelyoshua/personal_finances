package models

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mickaelyoshua/personal-finances/db/sqlc"
)

type SQLAgent struct {
	Conn    *pgx.Conn
	Queries *sqlc.Queries
}

func NewAgent(ctx context.Context, databaseURL string) (*SQLAgent, error) {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	queries := sqlc.New(conn)

	return &SQLAgent{
		Conn:    conn,
		Queries: queries,
	}, nil
}