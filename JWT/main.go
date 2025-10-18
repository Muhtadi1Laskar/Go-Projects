package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	JWT_SECRET = "This_is_a_secret_KEY"
	ALGORTIHM = "HS256"
)

func base64urlEncode(data []byte) string {
	str := base64.URLEncoding.EncodeToString(data)
	return strings.TrimRight(str, "=")
}

func base64urlDecode(data string) ([]byte, error) {
	padding := len(data) % 4
	if padding > 0 {
		data += strings.Repeat("=", 4 - padding)
	}
	return base64.URLEncoding.DecodeString(data)
}



func main() {
	fmt.Println("Hello World")
}