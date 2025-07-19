package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"bip-39/Cipher"
	"bip-39/KeyOperations"
	"bip-39/byteOperations"
)

func getTxtFilePath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, "./bip-39-words.txt")
}

func main() {
	filePath := getTxtFilePath()
	mnemonic := byteoperations.GeneratePhrase(filePath)
	var mnemonicStr string = strings.Join(mnemonic, " ")

	const password string = "Hello890World"
	var seed []byte = byteoperations.GenerateSeed(mnemonicStr, password)
	masterKey, masterChain := keyoperations.GenerateMasterKey(seed)
	childIndex := uint32(0x80000000)

	childKey, childChain, err := keyoperations.DeriveHardenedChilds(masterKey, masterChain, childIndex)
	if err != nil {
		log.Fatalf("Child derivation failed: %v", err)
	}
	publicKey := keyoperations.PrivateKeyToPublicKey(childKey)
	address := keyoperations.GenerateP2PKeyAddress(publicKey)

	aesKey := ciphers.DeriveAESKey(password)
	encryptedPrivateKey, _ := ciphers.AESEncrypt(password, aesKey)
	decryptedPrivateKey, _ := ciphers.AESDecrypt(encryptedPrivateKey, aesKey)

	fmt.Println("🔐 Your 12-word mnemonic phrase:")
	fmt.Println(strings.Join(mnemonic, " ") + "\n")
	fmt.Printf("🔗 Child Chain Code: %x", childChain)
	fmt.Printf("\n🔐 Private Key: %x\n", childKey)
	fmt.Printf("📍 Public Key: %x\n", publicKey)
	fmt.Printf("📦 Address: %s\n", address)

	fmt.Printf("\nEncrypted Private Key: %s\n", encryptedPrivateKey)
	fmt.Printf("Decrypted Private Key: %s\n", decryptedPrivateKey)
}
