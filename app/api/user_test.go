package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/mickaelyoshua/personal_finances/db/mock"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRegister(t *testing.T) {
	user := randomUser(t)

	testCases := []struct {
		Name       string
		Email      string
		body       gin.H
		BuildStubs func(agent *mock.MockAgent)
		CheckResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name:  "OK",
			Email: user.Email,
			body: gin.H{
				"name":     user.Name,
				"email":    user.Email,
				"password": user.PasswordHash,
			},
			BuildStubs: func(agent *mock.MockAgent) {
				args := sqlc.CreateUserParams{
					Name:         util.RandomName(),
					Email:        util.RandomEmail(),
					PasswordHash: util.RandomPassword(),
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
			Email: user.Email,
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

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(data))
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			server := NewServer(agent)
			server.router.ServeHTTP(recorder, request)
			tc.CheckResponse(t, recorder)
		})
	}
}

func TestLogin(t *testing.T) {
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

	testCases := []struct {
		Name string
		body gin.H
		BuildStubs func(agent *mock.MockAgent)
		CheckResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "OK",
			body: gin.H{
				"email":    user.Email,
				"password": user.PasswordHash,
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

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(data))
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			server := NewServer(agent)
			server.router.ServeHTTP(recorder, request)
			tc.CheckResponse(t, recorder)
		})
	}
}

func randomUser(t *testing.T) sqlc.User {
	password := util.RandomPassword()
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	return sqlc.User{
		Name:         util.RandomName(),
		Email:        util.RandomEmail(),
		PasswordHash: hashedPassword,
	}
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