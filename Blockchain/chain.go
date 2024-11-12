package main

import (
	"fmt"
	"crypto/sha512"
	"encoding/hex"
)

func sha521(text string) string {
	byteMessage := []byte(text)
	hash := sha512.New()
	hash.Write(byteMessage)

	hashedBytes := hash.Sum(nil)
	encodedStr := hex.EncodeToString(hashedBytes)

	return encodedStr
}

func main() {
	fmt.Println(sha521("Hello World"))
}