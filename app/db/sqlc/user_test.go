package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mickaelyoshua/personal-finances/util"
	"github.com/stretchr/testify/require"
)

func deleteRandomUser(userID int32) error {
	return testQueries.HardDeleteUser(context.Background(), userID)
}

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
	// delete the user after test
	defer func() {
		err = deleteRandomUser(user.ID)
		require.NoError(t, err)
	}()

	require.NoError(t, err)
	require.NotEmpty(t, user)
}

func TestGetUserById(t *testing.T) {
	user1, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user1)
	// delete the user after test
	defer func() {
		err = deleteRandomUser(user1.ID)
		require.NoError(t, err)
	}()

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
	// delete the user after test
	defer func() {
		err = deleteRandomUser(user1.ID)
		require.NoError(t, err)
	}()

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
	var usersID []int32
	numberOfUsers := 10
	for range numberOfUsers {
		user, err := createRandomUser()
		require.NoError(t, err)
		usersID = append(usersID, user.ID)
	}
	// delete users after test
	defer func() {
		for _, userID := range usersID {
			err := deleteRandomUser(userID)
			require.NoError(t, err)
		}
	}()

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

func TestGetAllUsersWithDeleted(t *testing.T) {
	var usersID []int32
	numberOfUsers := 5
	for range numberOfUsers {
		user, err := createRandomUser()
		require.NoError(t, err)
		usersID = append(usersID, user.ID)
	}
	// delete users after test
	defer func() {
		for _, userID := range usersID {
			err := deleteRandomUser(userID)
			require.NoError(t, err)
		}
	}()

	users, err := testQueries.GetAllUsersWithDeleted(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, users)
	require.GreaterOrEqual(t, len(users), numberOfUsers)

	for _, user := range users {
		require.NotEmpty(t, user)

		// Ensure that deleted_at can be nil or not nil
		if user.DeletedAt.Valid {
			require.NotEmpty(t, user.DeletedAt)
		} else {
			require.Empty(t, user.DeletedAt)
		}
	}
}

func TestUpdateUser(t *testing.T) {
	user1, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user1)
	// delete the user after test
	defer func() {
		err = deleteRandomUser(user1.ID)
		require.NoError(t, err)
	}()

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

	require.NotEqual(t, user1.Name, user2.Name) // Name should change
	require.NotEqual(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time) // UpdatedAt should change
	require.WithinDuration(t, updateTime, user2.UpdatedAt.Time, 2*time.Second) // UpdatedAt should be close to now

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, args.Name, user2.Name)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user1)
	// delete the user after test
	defer func() {
		err = deleteRandomUser(user1.ID)
		require.NoError(t, err)
	}()

	err = testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestRestoreUser(t *testing.T) {
	user1, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user1)
	// delete the user after test
	defer func() {
		err = deleteRandomUser(user1.ID)
		require.NoError(t, err)
	}()

	err = testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.RestoreUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Empty(t, user2.DeletedAt)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
	require.WithinDuration(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time, time.Second)
}

func TestHardDeleteUser(t *testing.T) {
	user1, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	err = testQueries.HardDeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, user2)
}