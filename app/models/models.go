package models

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mickaelyoshua/personal-finances/db/sqlc"
	"github.com/mickaelyoshua/personal-finances/util"
)

type SQLAgent struct {
	Conn    *pgx.Conn
	Queries *sqlc.Queries
}

func GetSQLAgent(ctx context.Context) (*SQLAgent, error) {
	databaseURL, err := util.GetDatabaseURL()
	if err != nil {
		return nil, err
	}

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