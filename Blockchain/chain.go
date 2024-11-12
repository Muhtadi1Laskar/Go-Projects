package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
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
	newBlock := &Block{
		index:        len(c.chain),
		timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		proof:        proof,
		previousHash: previousHash,
	}
	c.chain = append(c.chain, newBlock)

	return newBlock
}

func (c *Chain) getPrevBlock() *Block {
	return c.chain[len(c.chain)-1]
}

func (c *Chain) hash(block *Block) string {
	record := strconv.Itoa(block.index) + block.timestamp + strconv.Itoa(block.proof) + block.previousHash
	hash := sha521(record)
	return hash
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
	chain := &Chain{}
	block := chain.createBlock(12345, "0")

	fmt.Printf("%+v\n", block)
}
