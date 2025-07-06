package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
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
	validators map[string]int
}

func newBlockChain() *Chain {
	c := &Chain{validators: make(map[string]int)}
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
	chain.validators[name] = stake
}

func (chain *Chain) selectValidator() string {
	var (
		names []string
		stakes []int
		total int
	)

	for name, stake := range chain.validators {
		names = append(names, name)
		stakes = append(stakes, stake)
		total += stake
	}

	if total == 0 {
		return "Network"
	}

	r := rand.Intn(int(total))
	cumulative := 0
	for i, w := range stakes {
		cumulative += w
		if r < cumulative {
			return names[i]
		}
	}

	return " "
}

func (chain *Chain) addBlock(data string) *Block {
	validator := chain.selectValidator()
	lastBlock := chain.chain[len(chain.chain)-1]
	newBlock := &Block{
		index: len(chain.chain),
		timeStamp: time.Now().Format(time.RFC3339),
		data: data,
		previousHash: lastBlock.hash,
		validator: validator,
	}
	newBlock.hash = chain.hashBlock(newBlock)

	chain.chain = append(chain.chain, newBlock)

	fmt.Println("✅ Block added by: ", validator)

	return newBlock
}

func (chain *Chain) hashBlock(block *Block) string {
	record := strconv.Itoa(block.index) + block.data + block.timeStamp + block.previousHash
	hash := computeHash(record)
	return hash
}

func (chain *Chain) print() string {
	var result string
	for _, value := range chain.chain {
		hashPrefix := value.hash
		if len(hashPrefix) > 10 {
			hashPrefix = hashPrefix[:10]
		}
		result += fmt.Sprintf("Block #%d | Validator: %s | Hash: %s...\n", 
			value.index,
			value.validator, 
			hashPrefix)
	}
	return result
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

	blockChain.addValidators("Saitama", 70)
	blockChain.addValidators("Genos", 60)
	blockChain.addValidators("King", 20)

	for i := range 10 {
		blockChain.addBlock("Transcations: " + strconv.Itoa(i))
	}

	fmt.Print("\n📦 Blockchain State:\n")
	fmt.Println(blockChain.print())
}
