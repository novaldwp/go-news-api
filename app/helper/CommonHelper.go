package helper

import (
	"math/rand"
)

// random string with integer generator for slug with custom length character.
// e.g => randStrInt(4); res => i3C9, randStrInt(5); res => h9C78
func randStrInt(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
