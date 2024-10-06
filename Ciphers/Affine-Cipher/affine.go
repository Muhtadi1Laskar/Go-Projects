package main

import (
	"fmt"
	"unicode"
)

func affineEncrypt(text string, a, b int) string {
	encrypted := []rune{}

	for _, r := range text {
		if unicode.IsLetter(r) {
			offset := 'A'
			if unicode.IsLower(r) {
				offset = 'a'
			}

			x := int(r - offset)
			enc := (a*x + b) % 26
			encrypted = append(encrypted, offset+rune(enc))
		} else {
			encrypted = append(encrypted, r)
		}
	}

	return string(encrypted)
}

func main() {
	a := 5
	b := 8
	text := "Affine Cipher Example"

	encrypted := affineEncrypt(text, a, b)
	fmt.Println("Encrypted:", encrypted)
}