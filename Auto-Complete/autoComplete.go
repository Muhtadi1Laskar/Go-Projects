package main

import (
	"fmt"
	"strings"
)

type Node struct {
	left  *Node
	right *Node
	value string
}

type Tree struct {
	root *Node
}

func new_tree() *Tree {
	return &Tree{root: nil}
}

func (bst *Tree) insert(data string) {
	if bst.root == nil {
		bst.root = &Node{value: data}
	} else {
		bst.insertHelper(bst.root, data)
	}
}

func (bst *Tree) insertHelper(currentNode *Node, data string) {
	if currentNode.value > data {
		if currentNode.left == nil {
			currentNode.left = &Node{value: data}
		} else {
			bst.insertHelper(currentNode.left, data)
		}
	} else {
		if currentNode.right == nil {
			currentNode.right = &Node{value: data}
		} else {
			bst.insertHelper(currentNode.right, data)
		}
	}
}

func (bst *Tree) autoComplete(prefix string) []string {
	var result []string
	var traverse func(n *Node)
	traverse = func(n *Node) {
		if n == nil {
			return
		}

		if strings.HasPrefix(n.value, prefix) {
			traverse(n.left)
			result = append(result, n.value)
			traverse(n.right)
		} else if n.value > prefix {
			traverse(n.left)
		} else {
			traverse(n.right)
		}
	}
	traverse(bst.root)
	return result
}

func (bst *Tree) iterativeDFS() []string {
	var result []string
	var stack []*Node
	var currentNode = bst.root

	for currentNode != nil || len(stack) > 0 {
		for currentNode != nil {
			stack = append(stack, currentNode)
			currentNode = currentNode.left
		}
		currentNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, currentNode.value)

		currentNode = currentNode.right
	}
	return result
}

func main() {
	var tree = new_tree()
	var words []string = []string{"apple", "app", "apricot", "xoy", "banana", "ball", "cat", "a", "czech", "carrot"}

	for _, word := range words {
		tree.insert(word)
	}

	fmt.Println(tree.iterativeDFS())

	fmt.Println(tree.autoComplete("c"))
}
