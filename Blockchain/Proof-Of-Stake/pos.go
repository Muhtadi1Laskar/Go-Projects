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

func newBlockChain() *Chain {
	c := &Chain{validator: make(map[string]int)}
	c.chain = append(c.chain, c.genesisBlock())
	c.chain = append(c.chain, c.genesisBlock())
	return c
}

func (chain *Chain) genesisBlock() *Block {
	var block *Block = &Block{
		index:        0,
		timeStamp:    time.Now().Format(time.RFC3339),
		data:         "Genesis Block",
		previousHash: "0",
		validator:    "Network",
	}
	block.hash = chain.hashBlock(block)

	return block
}

func (chain *Chain) addValidators(name string, stake int) {
	chain.validator[name] = stake
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
	blockChain := newBlockChain()

	blockChain.addValidators("Saitama", 20)
	blockChain.addValidators("Genos", 30)
	blockChain.addValidators("King", 60)

	fmt.Println(blockChain.validator)

	for _, value := range blockChain.chain {
		fmt.Printf("%+v\n\n", value)
	}
}
