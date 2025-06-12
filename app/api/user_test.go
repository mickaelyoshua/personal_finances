package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mickaelyoshua/personal_finances/db/mock"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRegister(t *testing.T) {
	testCases := []struct {
		
	}{

	}


	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	agent := mock.NewMockAgent(ctrl)

	args := sqlc.CreateUserParams{
		Name:         util.RandomName(),
		Email:        util.RandomEmail(),
		PasswordHash: util.RandomPassword(),
	}
	user := sqlc.User{
		Name:         args.Name,
		Email:        args.Email,
		PasswordHash: args.PasswordHash,
	} 

	// Build stubs
	agent.EXPECT().
		CreateUser(gomock.Any(), gomock.Eq(args)).
		Times(1).
		Return(user, nil)
	
	// Start the server and send request
	server := NewServer(agent)
	recorder := httptest.NewRecorder()
	url := "/register"
	request, err := http.NewRequest(http.MethodPost, url, userToReader(user))
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	
	// Check the response
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchUser(t, recorder.Body, user)
}

func requireBodyMatchUser(t *testing.T, body io.Reader, user sqlc.User) {
	gotUser, err := readerToUser(body)
	require.NoError(t, err)

	require.Equal(t, user.Name, gotUser.Name)
	require.Equal(t, user.Email, gotUser.Email)
	require.Equal(t, user.PasswordHash, gotUser.PasswordHash)
}

func userToReader(user sqlc.User) io.Reader {
	str := `{"name":"` + user.Name + `","email":"` + user.Email + `","password_hash":"` + user.PasswordHash + `"}`
	return strings.NewReader(str)
}
func readerToUser(body io.Reader) (sqlc.User, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return sqlc.User{}, err
	}

	var user sqlc.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}
