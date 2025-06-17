package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/mickaelyoshua/personal_finances/db/mock"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRegister(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		Name       string
		FormData url.Values
		BuildStubs func(agent *mock.MockAgent)
		CheckResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name:  "OK",
			FormData: url.Values{
				"name":     {user.Name},
				"email":    {user.Email},
				"password": {password},
			},
			BuildStubs: func(agent *mock.MockAgent) {
				args := sqlc.CreateUserParams{
					Name:         user.Name,
					Email:        user.Email,
					PasswordHash: user.PasswordHash,
				}
				agent.EXPECT().
					CreateUser(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(user, nil)
			},
			CheckResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusSeeOther, recorder.Code)
			},
		},
		{
			Name:  "InternalServerError",
			BuildStubs: func(agent *mock.MockAgent) {
				agent.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sqlc.User{}, pgx.ErrNoRows)
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

			request, err := http.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(tc.FormData.Encode()))
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			recorder := httptest.NewRecorder()
			server := NewServer(agent)
			server.router.ServeHTTP(recorder, request)
			tc.CheckResponse(t, recorder)
		})
	}
}

func TestLogin(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		Name string
		FormData url.Values
		BuildStubs func(agent *mock.MockAgent)
		CheckResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "OK",
			FormData: url.Values{
				"email":  {user.Email},
				"password": {password},
			},
			BuildStubs: func(agent *mock.MockAgent) {
				agent.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
			},
			CheckResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusSeeOther, recorder.Code)
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

			request, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(tc.FormData.Encode()))
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			recorder := httptest.NewRecorder()
			server := NewServer(agent)
			server.router.ServeHTTP(recorder, request)
			tc.CheckResponse(t, recorder)
		})
	}
}

func randomUser(t *testing.T) (sqlc.User, string) {
	password := util.RandomPassword()
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	return sqlc.User{
		Name:         util.RandomName(),
		Email:        util.RandomEmail(),
		PasswordHash: hashedPassword,
	}, password
}

//func requireBodyMatchUser(t *testing.T, body io.Reader, user sqlc.User) {
//	data, err := io.ReadAll(body)
//	require.NoError(t, err)
//
//	var gotUser sqlc.User
//	err = json.Unmarshal(data, &gotUser)
//	require.NoError(t, err)
//
//	require.Equal(t, user.Name, gotUser.Name)
//	require.Equal(t, user.Email, gotUser.Email)
//	require.Equal(t, user.PasswordHash, gotUser.PasswordHash)
//}