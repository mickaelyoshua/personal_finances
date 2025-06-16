package sqlc

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/mickaelyoshua/personal_finances/util"
)

// testQueries is a global variable that holds the instance of Queries.
// Will be used across multiple test functions.
var testQueries *Queries

// testUser is a global variable that holds a test user created in TestMain.
// This user will be used in tests that require a user context.
var testUser User

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	config, err := util.LoadConfig("../../../")
	if err != nil {
		log.Fatal("Failed to load configuration: " + err.Error())
	}

	conn, err := pgx.Connect(context.Background(), config.DatabaseURL)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	testQueries = New(conn)

	testUser, err = testQueries.CreateUser(context.Background(), CreateUserParams{
		Name:         "Test User",
		Email:        "testuser@example.com",
		PasswordHash: "hashedpassword",
	})
	if err != nil {
		log.Fatalf("cannot create test user: %v", err)
	}
	log.Printf("Test user created with ID: %d", testUser.ID)

	code := m.Run()
	// Run after all tests are done, after the m.Run() call to ensure the test user is deleted after tests
	if err := testQueries.HardDeleteUser(context.Background(), testUser.ID); err != nil {
		log.Fatalf("cannot delete test user: %v", err)
	}
	log.Printf("Test user with ID %d deleted", testUser.ID)

	os.Exit(code)
}