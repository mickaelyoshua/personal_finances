package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/mickaelyoshua/personal_finances/db/mock"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type eqCreateUserParamsMatcher struct {
	Args sqlc.CreateUserParams
	Password string
}

// This is to match the interface implementation of gomock.Eq()
// a Matcher with two functions: Matches and String
func (e eqCreateUserParamsMatcher) Matches(x any) bool {
	params, ok := x.(sqlc.CreateUserParams)
	if !ok {
		return false
	}
	
	err := util.CompareHashPassword(params.PasswordHash, e.Password)
	if err != nil {
		return false
	}

	e.Args.PasswordHash = params.PasswordHash
	return reflect.DeepEqual(e.Args, params)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches args %v with password %v", e.Args, e.Password)
}

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
				paramsMatch := eqCreateUserParamsMatcher{
						Args:     args,
						Password: password,
					}
				agent.EXPECT().
					CreateUser(gomock.Any(), paramsMatch).
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
			server := newTestServer(t, agent)
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
			server := newTestServer(t, agent)
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