package main

import (
	"fmt"
)

func atbashCipher(data string) string {
	var result []rune

	for _, char := range data {
		if char >= 'A' && char <= 'Z' {
			result = append(result, 'Z' - char + 'A')
		} else if char >= 'a' && char <= 'z' {
			result = append(result, 'z' - char + 'a')
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

func main() {
	plaintext := "hello world"
    ciphertext := atbashCipher(plaintext)
	dicipherText := atbashCipher(ciphertext)

    fmt.Println("Cipher Text: ", ciphertext)
	fmt.Println("Decrypted Text: ", dicipherText)
}