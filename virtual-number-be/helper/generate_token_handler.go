package helper

import (
	"math/rand"
)

func GenerateToken(length int) string {
	digits := "0123456789"
	token := ""
	for i := 0; i < length; i++ {
		token += string(digits[rand.Intn(len(digits))])
	}
	return token
}
