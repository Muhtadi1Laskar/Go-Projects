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

func (list *LinkedList) Print() {
	if list.isEmpty() {
		return 
	}
	currentNode := list.head

	for i := 1; currentNode != nil; i++ {
		fmt.Printf("%d: %s\n", i, currentNode.value)
		currentNode = currentNode.next
	}
}

func main() {
	tasks := NewLinkedList()

	tasks.Append("C")
	tasks.Append("JavaScript")
	tasks.Append("Python")
	tasks.Append("Go")
	tasks.Append("Julia")
	tasks.Append("Odin")
	tasks.Append("Haskell")

	tasks.Print()
}