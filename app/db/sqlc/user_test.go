package sqlc

import (
	"context"
	"testing"

	"github.com/mickaelyoshua/personal-finances/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	args := CreateUserParams{
		Name:         util.RandomName(),
		Email:        util.RandomEmail(),
		PasswordHash: util.RandomPassword(),
	}
	user, err := testQueries.CreateUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Name, user.Name)
	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
}