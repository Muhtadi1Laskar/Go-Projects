package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readWordList(filePath string) []string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read wordlist: %v", err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func hashBytes(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)

	hashedBytes := hash.Sum(nil)
	return hashedBytes
}

func main() {
	var total int = 16

	var randomBytes []byte = make([]byte, total)

	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Error while generating random number", err)
	}

	hashedBytes := hashBytes(randomBytes)
	firstByte := randomBytes[0]
	first4Bits := firstByte >> 4

	fmt.Printf("Generated Random Bytes (hex): %x\n", randomBytes)
	fmt.Printf("Generated Random Bytes (raw): %v\n", randomBytes)
	fmt.Printf("Generated Random Bytes (hashed raw): %v\n", hashedBytes)
	fmt.Println(hex.EncodeToString(hashedBytes))
	fmt.Println(first4Bits)
}