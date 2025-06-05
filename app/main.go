package main

import (
	"context"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/mickaelyoshua/personal-finances/personal_finances"
)

func run() error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=postgres password=postgres dbname=personal_finances host=localhost port=5432 sslmode=verify-full")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := personal_finances.New(conn)

	insertedUser, err := queries.CreateUser(ctx, personal_finances.CreateUserParams{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	})
	if err != nil {
		return err
	}
	log.Printf("Inserted User: %+v\n", insertedUser)

	user, err := queries.GetUser(ctx, insertedUser.ID)
	if err != nil {
		return err
	}
	log.Printf("Retrieved User: %+v\n", user)

	if reflect.DeepEqual(insertedUser, user) {
		log.Println("Inserted user and retrieved user are identical.")
	} else {
		log.Println("Inserted user and retrieved user differ.")
	}
	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}