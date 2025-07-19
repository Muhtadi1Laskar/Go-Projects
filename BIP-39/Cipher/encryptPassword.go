package ciphers

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func DeriveAESKey(password string) string {
	salt := []byte("A_fucking_salt")
	aesKey := pbkdf2.Key([]byte(password), salt, 2048, 32, sha256.New)
	return string(aesKey)
}