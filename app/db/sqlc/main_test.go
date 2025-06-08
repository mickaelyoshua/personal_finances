package sqlc

import (
	"log"
	"os"
	"testing"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("cannot load .env file: %v", err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}