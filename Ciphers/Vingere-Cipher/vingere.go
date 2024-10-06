package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const PATH string = "C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects/Ciphers/Vingere-Cipher/text.txt"

func readFile() string {
	var builder strings.Builder

	file, err := os.Open(PATH)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while reading file: %v", err)
	}

	return builder.String()
}

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
	plainText := readFile()
	key := "KEY"

	encrypted := vingereCipher(plainText, key, true)
	decrypted := vingereCipher(encrypted, key, false)

	fmt.Println(encrypted)
	fmt.Println(decrypted)
}