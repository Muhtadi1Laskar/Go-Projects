package main

import "fmt"

func CaesarCipher(data string, key int) string {
	var result []rune

	for _, char := range data {
		if char >= 'A' && char <= 'Z' {
			result = append(result, rune('A'+(char-'A'+rune(key))%26))
		} else if char >= 'a' && char <= 'z' {
			result = append(result, rune('a'+(char-'a'+rune(key))%26))
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

func main() {
	var message string = "This is a message"
	var shift int = 2
	var cipherMessage string = CaesarCipher(message, shift)
	var decipherMessage string = CaesarCipher(cipherMessage, -shift)

	fmt.Println(cipherMessage)
	fmt.Println(decipherMessage)
}