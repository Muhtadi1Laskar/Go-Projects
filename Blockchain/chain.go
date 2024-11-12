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

type Response struct {
	message      string
	index        int
	timestamp    string
	proof        int
	previousHash string
}

type Chain struct {
	chain []*Block
}

func NewBlockChain() *Chain {
	c := &Chain{}
	c.createBlock(1, "0")
	return c
}

func (c *Chain) createBlock(proof int, previousHash string) *Block {
	newBlock := &Block{
		index:        len(c.chain) + 1,
		timestamp:    time.Now().Format(time.RFC3339),
		proof:        proof,
		previousHash: previousHash,
	}
	c.chain = append(c.chain, newBlock)

	return newBlock
}

func (c *Chain) getPrevBlock() *Block {
	return c.chain[len(c.chain)-1]
}

func (c *Chain) pow(prevProof int) int {
	var newProof int = 1
	var checkProof bool = false

	for !checkProof {
		num := (newProof * newProof) - (prevProof * prevProof)
		hashOperation := sha521(strconv.Itoa(num))

		if hashOperation[:4] == "0000" {
			checkProof = true
		} else {
			newProof++
		}
	}
	return newProof
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

func (chain *Chain) mineBlock() *Response {
	previousBlock := chain.getPrevBlock()
	previousProof := previousBlock.proof
	proof := chain.pow(previousProof)
	previousHash := chain.hash(previousBlock)
	block := chain.createBlock(proof, previousHash)

	return &Response{
		message:      "Congragulations you just mined a block",
		index:        block.index,
		timestamp:    block.timestamp,
		proof:        block.proof,
		previousHash: block.previousHash,
	}
}

func main() {
	chain := NewBlockChain()
	for i := 0; i < 10; i++ {
		block := chain.mineBlock()

		fmt.Printf(" Message: %s\n Index: %d\n Timestamp: %s\n Proof: %d\n Previous Hash: %s\n", block.message, block.index, block.timestamp, block.proof, block.message)
	}
}
