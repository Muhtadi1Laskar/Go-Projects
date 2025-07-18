package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"bip-39/KeyOperations"
	"golang.org/x/crypto/pbkdf2"
)

// Reads the BIP-39 English wordlist from a file
func readWordList(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath) // ioutil.ReadFile is deprecated
	if err != nil {
		return nil, fmt.Errorf("failed to read wordlist: %w", err)
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	return lines, nil
}

func bytesToBits(data []byte) string {
	bits := make([]byte, 0, len(data)*8)
	for _, b := range data {
		bits = append(bits, bitToChar(b>>7))
		bits = append(bits, bitToChar(b>>6))
		bits = append(bits, bitToChar(b>>5))
		bits = append(bits, bitToChar(b>>4))
		bits = append(bits, bitToChar(b>>3))
		bits = append(bits, bitToChar(b>>2))
		bits = append(bits, bitToChar(b>>1))
		bits = append(bits, bitToChar(b))
	}
	return string(bits)
}

func bitToChar(b byte) byte {
	return '0' + (b & 1)
}

func getTxtFilePath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, "./bip-39-words.txt")
}

func generateSeed(mnemonic, password string) []byte {
	salt := "mnemonic" + password
	iterations := 2048
	keyLen := 64
	passwordBytes := []byte(mnemonic)
	saltBytes := []byte(salt)

	return pbkdf2.Key(passwordBytes, saltBytes, iterations, keyLen, sha512.New)
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

	wordlist, _ := readWordList(filePath)

	// Step 1: Generate 128-bit entropy (16 bytes for 12 words)
	entropy := make([]byte, 16)
	_, err := rand.Read(entropy)
	if err != nil {
		log.Fatalf("Entropy generation failed: %v", err)
	}

	// Step 2: Compute checksum (first N bits of SHA256(entropy), N = ENT / 32)
	hash := sha256.Sum256(entropy)
	entropyBits := bytesToBits(entropy)
	checksumBits := bytesToBits([]byte{hash[0]})[:4] // 128 bits → 4-bit checksum

	// Step 3: Concatenate entropy + checksum
	fullBits := entropyBits + checksumBits

	// Step 4: Split into 11-bit chunks and map to words
	var mnemonic []string
	for i := 0; i < len(fullBits); i += 11 {
		chunk := fullBits[i : i+11]
		index, err := strconv.ParseInt(chunk, 2, 64)
		if err != nil {
			log.Fatalf("Bit parsing failed: %v", err)
		}
		mnemonic = append(mnemonic, wordlist[index])
	}

	fmt.Println("🔐 Your 12-word mnemonic phrase:")
	fmt.Println(strings.Join(mnemonic, " ") + "\n")

	var seed []byte = generateSeed(strings.Join(mnemonic, " "), "hello90world")
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
