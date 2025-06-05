package main

import (
	"context"
	"log"
	"os"
	"io"
	"reflect"

	"github.com/jackc/pgx/v5"

	"github.com/mickaelyoshua/personal-finances/sqlc_generated"
)

func GetQueryFromFile() string {
	file, err := os.Open("./schema.sql")
	if err != nil {
		log.Fatalf("Failed to open schema.sql: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read schema.sql: %v", err)
	}

	return string(content)
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=postgres password=postgres dbname=personal_finance host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	schema := GetQueryFromFile()
	_, err = conn.Exec(ctx, schema)
	if err != nil {
		log.Fatalf("Failed to execute schema.sql: %v", err)
	}

	queries := sqlc_generated.New(conn)

	insertedUser, err := queries.CreateUser(ctx, sqlc_generated.CreateUserParams{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		PasswordHash: "$2a$10$EIX/1z5Z9Q8b1e5f3c5e6O0d7F4h5k1Z9Q8b1e5f3c5e6O0d7F4h5k1", // Example hash
	})
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("\n\nHI\n\n")
	log.Printf("Inserted User: %+v\n", insertedUser)

	user, err := queries.GetUser(ctx, insertedUser.ID)
	if err != nil {
		log.Fatalf("Failed to retrieve user: %v", err)
	}
	log.Printf("Retrieved User: %+v\n", user)

	if reflect.DeepEqual(insertedUser, user) {
		log.Println("Inserted user and retrieved user are identical.")
	} else {
		log.Println("Inserted user and retrieved user differ.")
	}
}