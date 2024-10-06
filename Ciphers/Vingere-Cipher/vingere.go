package main

import (
	"fmt"
	"unicode"
)

func vingereCipher(data string, key string, encrypt bool) string {
	keyRunes := []rune(key)
	textRunes := []rune(data)
	result := make([]rune, len(textRunes))

	for i, r := range textRunes {
		if !unicode.IsLetter(r) {
			result[i] = r
			continue
		}

		shift := keyRunes[i%len(keyRunes)]
		if unicode.IsUpper(shift) {
			shift -= 'A'
		} else {
			shift -= 'a'
		}

		if !encrypt {
			shift = -shift
		}

		if unicode.IsUpper(r) {
			result[i] = 'A' + (r - 'A' + shift + 26) % 26
		} else {
			result[i] = 'a' + (r - 'a' + shift + 26) % 26
		}
	}
	return string(result)
}

func main() {
	plainText := "Hello"
	key := "KEY"

	encrypted := vingereCipher(plainText, key, true)
	decrypted := vingereCipher(encrypted, key, false)

	fmt.Println(encrypted)
	fmt.Println(decrypted)
}