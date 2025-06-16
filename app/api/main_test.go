package api

import (
	"os"
	"testing"

	//"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	//gin.SetMode(gin.TestMode)
	code := m.Run()
	// Run after all tests are done, after the m.Run() call to ensure the test user is deleted after tests
	os.Exit(code)
}