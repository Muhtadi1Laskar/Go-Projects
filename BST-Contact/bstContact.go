package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	value map[string]string
}

type Tree struct {
	root *Node
}

func newTree() *Tree {
	return &Tree{root: nil}
}

func (bst *Tree) insert(data map[string]string) {
	if bst.root == nil {
		bst.root = &Node{ value: data }
	} else {
		bst.insertHelper(bst.root, data)
	}
}

func (bst *Tree) insertHelper(node *Node, data map[string]string) {
	if node.value["name"] < data["name"] {
		if node.left == nil {
			node.left = &Node{value: data}
		} else {
			bst.insertHelper(node.left, data)
		}
	} else {
		if node.right == nil {
			node.right = &Node{value: data}
		} else {
			bst.insertHelper(node.right, data)
		}
	}
}

func main() {
	var arr []map[string]string = []map[string]string{
		{
			"name":   "Luffy",
			"number": "01775900737",
		},
		{
			"name":   "Zoro",
			"number": "01338577464",
		},
		{
			"name":   "Sanji",
			"number": "55337866546",
		},
	}
	tree := newTree()

	for _, data := range arr {
		tree.insert(data)
	}

	fmt.Println(tree.root)
}
