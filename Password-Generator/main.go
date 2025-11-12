package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var PASSWORD_TYPES = map[string]string{
	"lower": "abcdefghijklmnopqrstuvwxyz",
	"upper": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"digits": "0123456789",
	"symbols": "!@#$%^&*()-_=+[]{}|;:,.<>?/",
}

func getRandomChar(chars string) string {
	return string(chars[rand.Intn(len(chars))])
}

func generatePassword(length int, types []string) string {
	if length < 4 {
		return "Password must be greater than 4"
	}

	var result []string
	for i := range types {
		c := PASSWORD_TYPES[types[i]]
		if len(c) == 0 {
			continue
		}
		result = append(result, getRandomChar(c))
	}

	var charPool string = ""
	for i := range types {
		charPool += PASSWORD_TYPES[types[i]]
	}

	for len(result) < length {
		result = append(result, getRandomChar(charPool))
	}

	for i := len(result)-1; i > 0; i-- {
		j := rand.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}

	return strings.Join(result, "")
}

func main() {
	var types []string = []string{"lower", "digits", "upper", "symbols"}

	fmt.Println(generatePassword(16, types))
}