package keyoperations

import (
	"crypto/hmac"
	"crypto/sha512"
)

func hmacSha512(seed []byte) []byte {
	secretKey := []byte("Bitcoin seed")

	h := hmac.New(sha512.New, secretKey)
	h.Write(seed)

	return h.Sum(nil)
}

func GenerateMasterKey(seed []byte) ([]byte, []byte) {
	keyedHash := hmacSha512(seed)
	IL := keyedHash[:32]
	IR := keyedHash[32:]

	return IL, IR
}