package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	// "errors"
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

func signToken(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64urlEncode(h.Sum(nil))
}

func createJWT(payload map[string]interface{}, secret string) (string, error) {
	header := map[string]string{
		"alg": ALGORTIHM,
		"typ": "JWT",
	}

	payload["exp"] = time.Now().Unix() + 36000

	headerJSON, _ := json.Marshal(header)
	payloadJSON, _ := json.Marshal(payload)

	encodedHeader := base64urlEncode(headerJSON)
	encodedPayload := base64urlEncode(payloadJSON)

	tokenParts := fmt.Sprintf("%s.%s", encodedHeader, encodedPayload)
	signature := signToken(tokenParts, secret)

	return fmt.Sprintf("%s.%s", tokenParts, signature), nil

}



func main() {
	data := map[string]interface{}{
		"name":       "Luffy",
		"id":         42343543405,
		"occupation": "Pirate",
	}

	token, _ := createJWT(data, JWT_SECRET)
	fmt.Println(token)
}