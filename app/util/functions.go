package util

import (
	"math/rand"
	"regexp"
	"time"
)

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}

func GenerateRandomID() string {
	rand.Seed(time.Now().UnixNano())

	idLength := 8

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	id := make([]byte, idLength)
	for i := range id {
		id[i] = chars[rand.Intn(len(chars))]
	}

	return string(id)
}
