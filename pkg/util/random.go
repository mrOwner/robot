package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabetdigit = "abcdefghijklmnopqrstuvwxyz0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max.
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// RandomInt generates a random integer between min and max.
func RandomInt32(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

// RandomFloat32 generates a random float berween min and max.
func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

// RandomString generates a random string of length n.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabetdigit)

	for i := 0; i < n; i++ {
		c := alphabetdigit[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomFIGI generates a random FIGI.
func RandomFIGI() string {
	return strings.ToUpper(RandomString(12))
}

// RandomYear generates a random year.
func RandomYear() int {
	return RandomInt(1987, 2040)
}
