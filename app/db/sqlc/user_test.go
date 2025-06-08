package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mickaelyoshua/personal-finances/util"
	"github.com/stretchr/testify/require"
)


func createRandomUser() (User, error) {
	args := CreateUserParams{
		Name:         util.RandomName(),
		Email:        util.RandomEmail(),
		PasswordHash: util.RandomPassword(),
	}
	return testQueries.CreateUser(context.Background(), args)
}

func TestCreateUser(t *testing.T) {
	user, err := createRandomUser()

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Name, user.Name)
	require.Equal(t, user.Email, user.Email)
	require.Equal(t, user.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
}

func TestGetUserById(t *testing.T) {
	user1, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
	require.WithinDuration(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time, time.Second)
}

func TestGetUserByEmail(t *testing.T) {
	user1, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
	require.WithinDuration(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time, time.Second)
}

func TestGetAllUsers(t *testing.T) {
	numberOfUsers := 10
	for i := 0; i < numberOfUsers; i++ {
		_, err := createRandomUser()
		require.NoError(t, err)
	}

	users, err := testQueries.GetAllUsers(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, users)
	require.GreaterOrEqual(t, len(users), numberOfUsers)

	for _, user := range users {
		require.NotEmpty(t, user)

		// Ensure that deleted_at is nil for all users
		require.Empty(t, user.DeletedAt)
	}

}

func TestUpdateUser(t *testing.T) {
	user1, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	args := UpdateUserParams{
		ID:           user1.ID,
		Name:         util.RandomName(),
		Email:        user1.Email, // Keeping the same email for update
		PasswordHash: user1.PasswordHash, // Keeping the same password for update
	}

	user2, err := testQueries.UpdateUser(context.Background(), args)
	updateTime := time.Now()

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, args.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)

	require.NotEqual(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time) // UpdatedAt should change
	require.WithinDuration(t, updateTime, user2.UpdatedAt.Time, 2*time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	err = testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, user2)
}