package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt generates a random integer between min and max
func RandomInt(min, max int32) int32 {
	return min + r.Int31n(max-min+1)
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

// RandomDate generates a random date within the last year
func RandomDate() pgtype.Date {
	var date pgtype.Date
	start := time.Now().AddDate(-1, 0, 0) // One year ago
	end := time.Now()                      // Now
	randomValue := r.Int63n(end.Sub(start).Nanoseconds())
	randomTime := start.Add(time.Duration(randomValue))

	date.Scan(randomTime)
	return date
}

func RandomCategoryID() pgtype.Int4 {
	var id pgtype.Int4
	id.Scan(RandomInt(1, 4))
	return id
}
// RandomFloat generates a random float64 between min and max
func RandomFloat(min, max float64) float64 {
	result := min + r.Float64()*(max-min)
	if result == 0 {
		return min
	}
	return result
}
// RandomAmount generates a random amount between 10.00 and 100.00
func RandomAmount() pgtype.Numeric {
	var amount pgtype.Numeric
	number := strconv.FormatFloat(RandomFloat(10.00, 100.00), 'f', 2, 64) // Random amount between 10.00 and 100.00
	amount.Scan(number)
	return amount
}

// RandomDescription generates a random description of length 20
func RandomDescription() pgtype.Text {
	var description pgtype.Text
	description.Scan(RandomString(20)) // Random description of length 20
	return description
}