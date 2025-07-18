package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"bip-39/KeyOperations"
	"bip-39/byteOperations"
)

func getTxtFilePath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, "./bip-39-words.txt")
}


func hmacSha512(seed []byte) []byte {
	secretKey := []byte("Bitcoin seed")

	h := hmac.New(sha512.New, secretKey)
	h.Write(seed)
	hmacSum := h.Sum(nil)

	return hmacSum
}

func generateMasterKey(seed []byte) ([]byte, []byte) {
	keyedHash := hmacSha512(seed)
	IL := keyedHash[:32]
	IR := keyedHash[32:]

	return IL, IR
}

func main() {
	filePath := getTxtFilePath()
	mnemonic := byteoperations.GeneratePhrase(filePath)

	fmt.Println("🔐 Your 12-word mnemonic phrase:")
	fmt.Println(strings.Join(mnemonic, " ") + "\n")

	var mnemonicStr string = strings.Join(mnemonic, " ")
	var seed []byte = byteoperations.GenerateSeed(mnemonicStr, "hello90world")
	masterKey, masterChain := generateMasterKey(seed)
	childIndex := uint32(0x80000000)

	childKey, childChain, err := keyoperations.DeriveHardenedChilds(masterKey, masterChain, childIndex)
	if err != nil {
		log.Fatalf("Child derivation failed: %v", err)
	}
	publicKey := keyoperations.PrivateKeyToPublicKey(childKey)
	address := keyoperations.GenerateP2PKeyAddress(publicKey)

	fmt.Printf("🔗 Child Chain Code: %x", childChain)
	fmt.Printf("\n🔐 Child Private Key: %x\n", childKey)
	fmt.Printf("📍 Public Key: %x\n", publicKey)
	fmt.Printf("📦 Address: %s\n", address)
}
