package main

import (
	"fmt"
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

func main() {
	fmt.Println("Hello World")
}
