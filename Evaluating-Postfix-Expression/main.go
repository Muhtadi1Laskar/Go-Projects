package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top    *Node
	bottom *Node
	len    int
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(item int) {
	newNode := &Node{ value: item }

	if stack.top == nil {
		stack.top = newNode
		stack.bottom = newNode
		stack.len++;
		return
	}
	newNode.next = stack.top
	stack.top = newNode
	stack.len++
	return
}

func (stack *Stack) Pop() int {
	if stack.top == nil {
		fmt.Println("The stack is  empty")
		return -1
	}
	item := stack.top
	stack.top = item.next
	stack.len--

	return item.value
}

func (stack *Stack) Log() []int {
	if stack.top == nil {
		return nil
	}
	currentNode := stack.top
	var result []int

	for currentNode != nil {
		result = append(result, currentNode.value)
		currentNode = currentNode.next
	}
	return result
}

func isNumber(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func evaluateExpression(exp string) int {
	expression := strings.Fields(exp)
	stack := NewStack()

	for _, token := range expression {
		if isNumber(token) {
			num, _ := strconv.Atoi(token)
			stack.Push(num)
		} else if token == "+" || token == "-" || token == "/" || token == "*" {
			itemTwo := stack.Pop()
			itemOne := stack.Pop()
			result := calculate(itemOne, itemTwo, token)
			stack.Push(result)
		}
	}
	result := stack.Log()

	if len(result) > 1 {
		fmt.Println("invalid expression: insufficient values for operation")
		return -1
	}
	return result[0]
}

func calculate(itemOne int, itemTwo int, operator string) int {
	var result int

	switch operator {
	case "+":
		result = itemOne + itemTwo
	case "-":
		result = itemOne - itemTwo
	case "/":
		result = itemOne / itemTwo
	case "*":
		result = itemOne * itemTwo
	}

	return result
}


func main() {
	fmt.Println(evaluateExpression("9 2 3 8 4 + * -"))
}


