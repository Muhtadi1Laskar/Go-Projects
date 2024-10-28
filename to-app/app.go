package main

import "fmt"

type Node struct {
	value string
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{ }
}

func (list *LinkedList) isEmpty() bool {
	return list.head == nil
}

func main() {
	fmt.Println("Hello World")
}