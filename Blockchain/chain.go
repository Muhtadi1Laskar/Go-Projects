package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

type Block struct {
	index        int
	timestamp    string
	proof        int
	previousHash string
}

type Chain struct {
	chain []*Block
}

func (c *Chain) createBlock(proof int, previousHash string) *Block {
	newBlock := &Block {
		index: len(c.chain),
		timestamp: "01/01/20002",
		proof: proof,
		previousHash: previousHash,
	}
	c.chain = append(c.chain, newBlock)

	return newBlock
}

func sha521(text string) string {
	byteMessage := []byte(text)
	hash := sha512.New()
	hash.Write(byteMessage)

	hashedBytes := hash.Sum(nil)
	encodedStr := hex.EncodeToString(hashedBytes)

	return encodedStr
}

func main() {
	fmt.Println(sha521("Hello World"))
}
