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

func (list *LinkedList) Append(value string) {
	newNode := &Node{ value: value }

	if list.isEmpty() {
		list.head = newNode
		list.tail = newNode
		list.length++
		return
	}
	list.tail.next = newNode
	list.tail = newNode
	list.length++
	return
}

func main() {
	fmt.Println("Hello World")
}