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
	if node.value["name"] > data["name"] {
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

func (bst *Tree) search(name string) string {
	var current_node = bst.root

	for current_node != nil {
		if current_node.value["name"] > name {
			current_node = current_node.left
		} else if current_node.value["name"] < name {
			current_node = current_node.right
		} else {
			var value, _ = current_node.value["number"]
			return value
		}
	}
	return "There is no contact with the name " + name
}

func (bst *Tree) traverse_in_order() {
	var walk func(node *Node)
	walk = func(node *Node) {
		if node != nil {
			walk(node.left)
			fmt.Println(node.value)
			walk(node.right)
		}
	}
	walk(bst.root)
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

	
	tree.insert(map[string]string{ "name": "Nami", "number": "23445422678" })
	tree.traverse_in_order()

	fmt.Println(tree.search("Nami"))
}
