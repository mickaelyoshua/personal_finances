package util

import (
	"math/rand"
	"time"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomName generates a random name of length 6
func RandomName() string {
	return RandomString(6)
}

// RandomEmail generates a random email address
func RandomEmail() string {
	return RandomString(6) + "@example.com"
}

// RandomPassword generates a random password of length 10
func RandomPassword() string {
	return RandomString(10)
}