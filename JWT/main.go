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
	ALGORTIHM  = "HS256"
)

func base64urlEncode(data []byte) string {
	str := base64.URLEncoding.EncodeToString(data)
	return strings.TrimRight(str, "=")
}

func base64urlDecode(data string) ([]byte, error) {
	padding := len(data) % 4
	if padding > 0 {
		data += strings.Repeat("=", 4-padding)
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

func verifyJWT(token, secret string) (map[string]interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	encodedHeader := parts[0]
	encodedPayload := parts[1]
	signature := parts[2]

	tokenParts := fmt.Sprintf("%s.%s", encodedHeader, encodedPayload)
	expectedSig := signToken(tokenParts, secret)

	if !hmac.Equal([]byte(signature), []byte(expectedSig)) {
		return nil, errors.New("signature verification failed")
	}

	payloadBytes, err := base64urlDecode(encodedPayload)
	if err != nil {
		return nil, err
	}

	var payload map[string]any
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, err
	}

	if exp, ok := payload["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, errors.New("token expired")
		}
	}

	return payload, nil
}

func main() {
	data := map[string]interface{}{
		"name":       "Luffy",
		"id":         42343543405,
		"occupation": "Pirate",
	}

	token, _ := createJWT(data, JWT_SECRET)
	fmt.Println(token + "\n")

	payload, err := verifyJWT(token, JWT_SECRET)
	if err != nil {
		fmt.Println("Error verifying token: ", err)
		return
	}

	for key, value := range payload {
		fmt.Println(key, value)
	}
}
