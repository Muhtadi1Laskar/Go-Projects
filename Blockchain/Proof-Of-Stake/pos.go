package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Block struct {
	index        int
	timeStamp    string
	data         string
	previousHash string
	validator    string
	hash         string
}

type Chain struct {
	chain     []*Block
	validator map[string]int
}

func (chain *Chain) hashBlock(block *Block) string {
	record := strconv.Itoa(block.index) + block.data + block.timeStamp + block.previousHash
	hash := computeHash(record)
	return hash
}

func computeHash(data string) string {
	byteMessage := []byte(data)
	hash := sha512.New()
	hash.Write(byteMessage)

	hashedBytes := hash.Sum(nil)
	encodedStr := hex.EncodeToString(hashedBytes)

	return encodedStr
}

func main() {
	fmt.Println("Hello World")
}
