package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(10)

	hashedPass, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPass)

	err = CompareHashPassword(hashedPass, password)
	require.NoError(t, err)

	wrongPass := RandomString(10)
	err = CompareHashPassword(hashedPass, wrongPass)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}