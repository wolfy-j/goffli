package utils

import (
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))
}

// RandString returns random string (lowercased) of desired length.
func RandString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}
