package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, agent sqlc.Agent) *Server {
	config := util.Config{
		TokenSymmetricKey: util.RandomString(32),
		AccessTokenDuration: 15 * time.Minute,
	}
	server, err := NewServer(config, agent)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	code := m.Run()
	// Run after all tests are done, after the m.Run() call to ensure the test user is deleted after tests
	os.Exit(code)
}