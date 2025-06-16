package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/mickaelyoshua/personal_finances/db/mock"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRegister(t *testing.T) {
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

	testCases := []struct{
		Name string
		Email string
		BuildStubs func(agent *mock.MockAgent)
		CheckResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "OK",
			Email: args.Email,
			BuildStubs: func(agent *mock.MockAgent) {
				agent.EXPECT().
					CreateUser(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(user, nil)
			},
			CheckResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusSeeOther, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			Name: "InternalServerError",
			Email: args.Email,
			BuildStubs: func(agent *mock.MockAgent) {
				agent.EXPECT().
					CreateUser(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(user, gomock.AssignableToTypeOf(errors.New("")))
			},
			CheckResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			agent := mock.NewMockAgent(ctrl)
			tc.BuildStubs(agent)

			urlValues := url.Values{
				"name":     {user.Name},
				"email":    {user.Email},
				"password": {user.PasswordHash},
			}
			request, err := http.NewRequest(http.MethodPost, "auth/register", userToReader(urlValues))
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			server := NewServer(agent)
			server.router.ServeHTTP(recorder, request)
			tc.CheckResponse(t, recorder)
		})
	}
}

func requireBodyMatchUser(t *testing.T, body io.Reader, user sqlc.User) {
	gotUser, err := readerToUser(body)
	require.NoError(t, err)

	require.Equal(t, user.Name, gotUser.Name)
	require.Equal(t, user.Email, gotUser.Email)
	require.Equal(t, user.PasswordHash, gotUser.PasswordHash)
}

func userToReader(urlValues url.Values) io.Reader {
	return strings.NewReader(urlValues.Encode())
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
