package byteoperations

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

)

func ReadWordList(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Unable to read the file: %v", err)
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	return lines, nil
}

func BytesToBits(data []byte) string {
	bits := make([]byte, 0, len(data)*8)
	for _, b := range data {
		bits = append(bits, bitsToChar(b>>7))
		bits = append(bits, bitsToChar(b>>6))
		bits = append(bits, bitsToChar(b>>5))
		bits = append(bits, bitsToChar(b>>4))
		bits = append(bits, bitsToChar(b>>3))
		bits = append(bits, bitsToChar(b>>2))
		bits = append(bits, bitsToChar(b>>1))
		bits = append(bits, bitsToChar(b))
	}
	return string(bits)
}

func bitsToChar(b byte) byte {
	return '0' + (b & 1)
}

func generateEntropy() ([]byte, error) {
	entropy := make([]byte, 16)
	_, err := rand.Read(entropy)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate entropy: %v\n", err)
	}
	return entropy, nil
}

func GeneratePhrase(filePath string) []string {
	wordList, _ := ReadWordList(filePath)
	entropy, _ := generateEntropy()
	
	hash := sha256.Sum256(entropy)
	entropyBits := BytesToBits(entropy)
	checkSumBits := BytesToBits([]byte{hash[0]})[:4]

	fullBits := entropyBits + checkSumBits

	var mnemonic []string
	for i := 0; i < len(fullBits); i += 11 {
		chunk := fullBits[i : i+11]
		index, err := strconv.ParseInt(chunk, 2, 64)
		if err != nil {
			log.Fatalf("Bit parsing failed: %v", err)
		}
		mnemonic = append(mnemonic, wordList[index])
	}
	return mnemonic
}
