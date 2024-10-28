package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) isEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) Append(value string) {
	newNode := &Node{value: value}

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

func (list *LinkedList) Insert(index int, value string) {
	if index <= 0 || index > list.length+1 {
		fmt.Printf("Index out of bound\n")
		return
	}
	newNode := &Node{value: value}

	if list.isEmpty() {
		list.head = newNode
		list.tail = newNode
		list.length++
		return
	}
	if index == 1 {
		newNode.next = list.head
		list.head = newNode
		list.length++
		return
	}
	previousNode := list.head

	for i := 1; i < index-1; i++ {
		previousNode = previousNode.next
	}
	currentNode := previousNode.next
	previousNode.next = newNode
	newNode.next = currentNode
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
	fmt.Println()
}

func main() {
	tasks := NewLinkedList()

	tasks.Append("Read fiction book")
	tasks.Append("Go thorough Learn Go course")
	tasks.Append("Practise Mathematics")
	tasks.Append("Execrise")
	tasks.Append("Check hacker news")
	tasks.Append("Go for a walk")
	tasks.Append("Go to supermarker")

	tasks.Print()
	tasks.Insert(5, "Read the book Grokking Algorithm")
	tasks.Print()
}
