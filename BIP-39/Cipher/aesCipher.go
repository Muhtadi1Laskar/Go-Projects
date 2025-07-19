package ciphers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func AESEncrypt(plainText, secretKey string) (string, error) {
	transformedKey, err := transformKey(secretKey)
	if err != nil {
		return "", err
	}
	aesBlock, err := aes.NewCipher(transformedKey)
	if err != nil {
		return "", fmt.Errorf("unable to create the instance of aes cipher: %s", err)
	}

	gcm, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return "", fmt.Errorf("unable to create the gcm: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", fmt.Errorf("unable to generate a random number: %v", err)
	}

	cipherByte := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return hex.EncodeToString(cipherByte), nil
}

func AESDecrypt(cipherTextHex, secretKey string) (string, error) {
	transformedKey, err := transformKey(secretKey)
	if err != nil {
		return "", err
	}

	cipherText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return "", fmt.Errorf("invalid hex-encoded cipherText: %v", err)
	}

	aesBlock, err := aes.NewCipher(transformedKey)
	if err != nil {
		return "", fmt.Errorf("unable to create the instance of aes cipher: %s", err)
	}

	gcm, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return "", fmt.Errorf("unable to create the gcm: %v", err)
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", fmt.Errorf("unable to decrypt the data: %v", err)
	}

	return string(plainText), nil
}

func transformKey(key string) ([]byte, error) {
	salt := []byte("Some-random-salt")
	iterations := 10000
	keyLength := 16

	secretKey := pbkdf2.Key([]byte(key), salt, iterations, keyLength, sha256.New)

	return secretKey, nil
}
